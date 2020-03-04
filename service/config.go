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
	"github.com/spf13/viper"
	"gopkg.in/natefinch/lumberjack.v2"
	"gopkg.in/yaml.v2"
)

var (
	conf models.Config
	db   *ent.Client
)

func initConf() {
	v := viper.New()
	v.SetConfigType("yaml")
	b, err := yaml.Marshal(models.Config{
		Salt: "",
		Env:  "dev",
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

	dbConf := conf.Db
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", dbConf.Host, dbConf.Port, dbConf.User, dbConf.Password, dbConf.Dbname)
	db, err = ent.Open("postgres", connStr, connOptions...)
	if err != nil {
		log.Fatal(err)
	}
	// run the auto migration tool.
	if err := db.Schema.Create(
		context.Background(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}

func initRole() {
	ctx := context.Background()
	exist, err := db.Role.Query().Where(role.ID(1)).Exist(ctx)
	if err != nil {
		log.Fatalf("init role error", err)
	}
	if exist {
		return
	}
	if _, err := db.Role.Create().SetName("默认角色").Save(ctx); err != nil {
		log.Fatalf("init role error", err)
	}
}

func GetConf() models.Config {
	return conf
}

func GetDB() *ent.Client {
	return db
}

func Init(e *echo.Echo) {
	initConf()
	if conf.Env == "dev" {
		e.Debug = true
		e.Logger.SetLevel(log.DEBUG)
		e.Logger.SetHeader("${time_rfc3339} ${level} ${line}")
	}
	e.Validator = models.NewValidator()
	initLogger(e)
	initDb(e)
	initRole()
}
