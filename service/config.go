package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/fitzix/assassin/consts"
	"github.com/fitzix/assassin/models"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
	"github.com/markbates/pkger"
	"github.com/minio/minio-go/v6"
	"github.com/rubenv/sql-migrate"
	"github.com/spf13/viper"
	"gopkg.in/natefinch/lumberjack.v2"
	"gopkg.in/yaml.v2"
)

var (
	conf models.Config
	db   *sqlx.DB
	s3   *minio.Client
)

func initConf() {
	v := viper.New()
	v.SetConfigType("yaml")
	b, err := yaml.Marshal(models.Config{
		Salt: "",
		Mod:  "dev",
		Db: models.Db{
			Host:     "127.0.0.1",
			Port:     5432,
			User:     "fitz",
			Password: "131833",
			Dbname:   "assassin-ent",
			SSL:      false,
		},
		Jwt: models.Jwt{
			Issuer:  "asn.xyz",
			Expires: 24 * 3,
			Secret:  "asn.io",
		},
		Encrypt: models.Encrypt{
			Key: "3C221351CA73FFA6",
		},
		S3: models.S3{
			Endpoint:        "play.min.io",
			AccessKeyID:     "Q3AM3UQ867SPQQA43P2F",
			SecretAccessKey: "zuf+tfteSlswRu7BJ86wekitnifILbZam1KYY3TG",
			UseSSL:          false,
			Bucket:          "assassin",
		},
	})

	if err != nil {
		log.Fatalf("maeshal default err: %s", err)
	}
	if err := v.MergeConfig(bytes.NewReader(b)); err != nil {
		log.Fatalf("merge default err: %s", err)
	}
	v.SetConfigFile("config.yml")
	if err := v.MergeInConfig(); err != nil {
		if _, ok := err.(viper.ConfigParseError); ok {
			log.Fatalf("merge default err: %s", err)
		}
	}
	// tell viper to overwrite env variables
	v.AutomaticEnv()
	v.SetEnvPrefix("ASN")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := v.Unmarshal(&conf); err != nil {
		log.Fatalf("unmarshal default err: %s", err)
	}
}

func initEcho(e *echo.Echo) {
	if conf.Mod == "dev" {
		e.Debug = true
		e.Logger.SetLevel(log.DEBUG)
		e.Logger.SetHeader("${time_rfc3339} ${level} ${prefix} ${short_file} ${line}")
	}
	e.Validator = models.NewValidator()
}

func initLogger(e *echo.Echo) {
	if e.Debug {
		return
	}

	hook := lumberjack.Logger{
		Filename: "logs/app.log",
		// 每个日志文件保存的最大尺寸 单位：M
		MaxSize: 100,
		// 日志文件最多保存多少个备份
		MaxBackups: 3,
		// 文件最多保存多少天
		MaxAge: 1,
		// 是否压缩
		Compress: true,
	}

	e.Logger.SetOutput(&hook)
}

func initDb(e *echo.Echo) {
	var err error
	connOption := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", conf.Host, conf.Port, conf.User, conf.Password, conf.Dbname)
	db, err = sqlx.Open("postgres", connOption)
	if err != nil {
		e.Logger.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		e.Logger.Fatal(err)
	}

	migrations := &migrate.HttpFileSystemMigrationSource{
		FileSystem: pkger.Dir("/migrations"),
	}

	if _, err := migrate.Exec(db.DB, "postgres", migrations, migrate.Up); err != nil {
		e.Logger.Fatalf("migrate err: %s", err)
	}
}

func initS3(e *echo.Echo) {
	var err error
	s3, err = minio.New(conf.Endpoint, conf.AccessKeyID, conf.SecretAccessKey, conf.UseSSL)
	if err != nil {
		e.Logger.Fatalf("connect minio file server err: %s", err)
	}
	// check bucket exist
	exist, err := s3.BucketExists(conf.Bucket)
	if err != nil {
		e.Logger.Fatalf("check minio bucket err: %s", err)
	}
	if exist {
		e.Logger.Info("minio bucket check ok")
		checkAndSetBucketPolicy(e)
		return
	}

	if err := s3.MakeBucket(conf.Bucket, "ap-east-1"); err != nil {
		e.Logger.Fatalf("create minio bucket err: %s", err)
	}

	setS3Policy(e, consts.S3PolicyAllowImageStatic)
}

// 检查图片资源是否公开
func checkAndSetBucketPolicy(e *echo.Echo) {
	p, err := s3.GetBucketPolicy(conf.Bucket)
	if err != nil {
		e.Logger.Fatalf("s3 get bucket policy err: %s", err)
		return
	}
	if p == "" {
		setS3Policy(e, consts.S3PolicyAllowImageStatic)
		return
	}

	var policy models.S3Policy
	if err := json.Unmarshal([]byte(p), &policy); err != nil {
		e.Logger.Fatalf("s3 parse bucket policy err: %s", err)
		return
	}
	if len(policy.Statement) > 0 {
		for _, v := range policy.Statement {
			if v.Sid == "AllowImageStatic" {
				e.Logger.Info("s3 bucket policy checked ok")
				return
			}
		}
	}

	var initPolicy models.S3Policy
	if err := json.Unmarshal([]byte(consts.S3PolicyAllowImageStatic), &initPolicy); err != nil {
		e.Logger.Fatalf("s3 parse init bucket policy err: %s", err)
		return
	}

	policy.Statement = append(policy.Statement, initPolicy.Statement[0])

	b, err := json.Marshal(&policy)
	if err != nil {
		e.Logger.Fatalf("s3 marshal new bucket policy err: %s", err)
		return
	}
	setS3Policy(e, string(b))
}

func setS3Policy(e *echo.Echo, policy string) {
	if err := s3.SetBucketPolicy(conf.Bucket, fmt.Sprintf(policy, conf.Bucket)); err != nil {
		e.Logger.Fatalf("set s3 bucket policy err: %s", err)
	}
	e.Logger.Printf("successfully set bucket policy %s", conf.Bucket)
}

func GetConf() models.Config {
	return conf
}

func GetDB() *sqlx.DB {
	return db
}

func GetS3() *minio.Client {
	return s3
}

func Init(e *echo.Echo) {
	initConf()
	initEcho(e)
	initLogger(e)
	initS3(e)
	initDb(e)
}
