package service

import (
	"bytes"
	"context"
	"fmt"
	"strings"

	"github.com/fitzix/assassin/ent"
	"github.com/fitzix/assassin/ent/migrate"
	"github.com/fitzix/assassin/ent/role"
	"github.com/fitzix/assassin/models"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
	"github.com/minio/minio-go/v6"
	"github.com/spf13/viper"
	"gopkg.in/natefinch/lumberjack.v2"
	"gopkg.in/yaml.v2"
)

var (
	conf models.Config
	db   *ent.Client
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
	var (
		err         error
		connOptions []ent.Option
	)

	if e.Debug {
		connOptions = append(connOptions, ent.Debug(), ent.Log(e.Logger.Info))
	}

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", conf.Host, conf.Port, conf.User, conf.Password, conf.Dbname)
	db, err = ent.Open("postgres", connStr, connOptions...)
	if err != nil {
		e.Logger.Fatal(err)
	}
	// run the auto migration tool.
	if err := db.Schema.Create(
		context.Background(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	); err != nil {
		e.Logger.Fatalf("failed creating schema resources: %v", err)
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
	// (ARN) Amazon 资源名称 唯一标识 AWS 资源
	// arn:partition:service:region:namespace:relative-id
	// 协议:分区:服务:区域:bucket名称:资源路径
	policy := `
		{
			"Version":"2012-10-17",
			"Statement":[
				{
					"Sid":"AllowImageStatic",
					"Action":"s3:GetObject",
					"Effect":"Allow",
					"Principal": "*",
					"Resource":[
						"arn:aws:s3:::%s/images/*"
					]
				}
			]
		}
	`
	if err := s3.SetBucketPolicy(conf.Bucket, fmt.Sprintf(policy, conf.Bucket)); err != nil {
		e.Logger.Fatalf("set s3 bucket policy err: %s", err)
	}
	e.Logger.Printf("Successfully created %s", conf.Bucket)
}

func checkAndSetBucketPolicy(e *echo.Echo) {
	// p, err := s3.GetBucketPolicy(conf.Bucket)
	// if err != nil {
	// 	e.Logger.Fatalf("s3 get bucket policy err: %s", err)
	// 	return
	// }
	// e.Logger.Fatalf("s3 get bucket policy err: %s", p)
}

func initRole(e *echo.Echo) {
	ctx := context.Background()
	exist, err := db.Role.Query().Where(role.ID(1)).Exist(ctx)
	if err != nil {
		e.Logger.Fatalf("init role error", err)
	}
	if exist {
		return
	}
	if _, err := db.Role.Create().SetName("默认角色").Save(ctx); err != nil {
		e.Logger.Fatalf("init role error", err)
	}
}

func GetConf() models.Config {
	return conf
}

func GetDB() *ent.Client {
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
	initRole(e)
}
