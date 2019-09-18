package service

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fitzix/assassin/db"
	"github.com/fitzix/assassin/models"
	"github.com/gin-gonic/gin"
	"github.com/google/go-github/v28/github"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/oauth2"
	"gopkg.in/natefinch/lumberjack.v2"
	"gopkg.in/yaml.v2"
)

var (
	appConf      models.Config
	zapLogger    *zap.Logger
	githubClient *GithubClient
	dbInstance   *gorm.DB
)

func InitProject() {
	initConf()
	initLogger()
	initGithubClient()
	initDb()
}

func initGithubClient() {
	conf := appConf.Github

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{
		AccessToken: conf.Token,
	})
	tc := oauth2.NewClient(ctx, ts)

	githubClient = &GithubClient{
		client: github.NewClient(tc),
		Github: conf,
		ctx:    ctx,
	}
}

func initConf() {
	v := viper.New()
	v.SetConfigType("yaml")
	b, err := yaml.Marshal(models.NewConfig())
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
		// dont return error if file is missing. overwrite file is optional
	}
	// tell viper to overwrite env variables
	v.AutomaticEnv()
	v.SetEnvPrefix("ASN")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := v.Unmarshal(&appConf); err != nil {
		log.Fatalf("unmarshal default err: %s", err)
	}
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
	encoder := zapcore.NewJSONEncoder(encoderConfig)

	var core zapcore.Core

	if gin.Mode() == gin.ReleaseMode {
		core = zapcore.NewCore(encoder, zapcore.AddSync(&hook), zap.InfoLevel)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
		core = zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zap.InfoLevel)
	}

	zapLogger = zap.New(core, zap.AddStacktrace(zapcore.ErrorLevel), zap.AddCaller())
}

func initDb() {
	conf := appConf.Db
	var err error
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", conf.Host, conf.Port, conf.User, conf.Password, conf.Dbname)
	dbInstance, err = gorm.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	if gin.Mode() != gin.ReleaseMode {
		dbInstance.LogMode(true)
	}
	dbInstance.SingularTable(true)
	db.MigrateDb(dbInstance.DB())
}

func GetConf() models.Config {
	return appConf
}

func GetLogger() *zap.Logger {
	return zapLogger
}

func GetGithubClient() *GithubClient {
	return githubClient
}

func GetDB() *gorm.DB {
	return dbInstance
}
