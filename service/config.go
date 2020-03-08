package service

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fitzix/assassin/consts"
	"github.com/fitzix/assassin/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/markbates/pkger"
	"github.com/minio/minio-go/v6"
	"github.com/rubenv/sql-migrate"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"gopkg.in/yaml.v3"
)

type asnLogger struct {
	*zap.SugaredLogger
}

var (
	conf   models.Config
	logger *asnLogger
	db     *gorm.DB
	s3     *minio.Client
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

func (l *asnLogger) Print(v ...interface{}) {
	l.Info(v...)
}

func initLogger() {
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

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	var core zapcore.Core

	if gin.Mode() == gin.ReleaseMode {
		core = zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderConfig),
			zapcore.AddSync(&hook),
			zap.InfoLevel,
		)
	} else {
		core = zapcore.NewCore(
			zapcore.NewConsoleEncoder(encoderConfig),
			zapcore.AddSync(os.Stdout),
			zap.InfoLevel,
		)
	}

	logger = &asnLogger{
		SugaredLogger: zap.New(core, zap.AddStacktrace(zapcore.ErrorLevel), zap.AddCaller()).Sugar(),
	}
}

func initDb() {
	var err error
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", conf.Host, conf.Port, conf.User, conf.Password, conf.Dbname)
	db, err = gorm.Open("postgres", connStr)
	if err != nil {
		logger.Fatal(err)
	}
	db.SetLogger(logger)
	db.LogMode(true)
	if gin.Mode() == gin.ReleaseMode {
		db.LogMode(false)
	}
	migrations := &migrate.HttpFileSystemMigrationSource{
		FileSystem: pkger.Dir("/migrations"),
	}
	if _, err := migrate.Exec(db.DB(), "postgres", migrations, migrate.Up); err != nil {
		logger.Fatalf("migrate err: %s", err)
	}
}

func initS3() {
	var err error
	s3, err = minio.New(conf.Endpoint, conf.AccessKeyID, conf.SecretAccessKey, conf.UseSSL)
	if err != nil {
		logger.Fatalf("connect minio file server err: %s", err)
	}
	// check bucket exist
	exist, err := s3.BucketExists(conf.Bucket)
	if err != nil {
		logger.Fatalf("check minio bucket err: %s", err)
	}
	if exist {
		logger.Info("minio bucket check ok")
		checkAndSetBucketPolicy()
		return
	}

	if err := s3.MakeBucket(conf.Bucket, "ap-east-1"); err != nil {
		logger.Fatalf("create minio bucket err: %s", err)
	}

	setS3Policy(consts.S3PolicyAllowImageStatic)
}

func GetConf() models.Config {
	return conf
}

func GetLogger() *zap.SugaredLogger {
	return logger.SugaredLogger
}

func GetDB() *gorm.DB {
	return db
}

func GetS3() *minio.Client {
	return s3
}

func Init() {
	initConf()
	initLogger()
	initS3()
	initDb()
}
