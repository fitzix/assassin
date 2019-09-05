package db

import (
	"fmt"
	"log"

	"github.com/fitzix/assassin/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/rubenv/sql-migrate"
)

var db *gorm.DB

func Init(conf models.Config) {
	var err error
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", conf.Db.Host, conf.Db.Port, conf.Db.User, conf.Db.Password, conf.Db.Dbname)
	db, err = gorm.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	if gin.Mode() != gin.ReleaseMode {
		db.LogMode(true)
	}
	db.SingularTable(true)
	migrateDb()
}

func migrateDb() {
	migrations := &migrate.FileMigrationSource{Dir: "db/migrations"}
	if _, err := migrate.Exec(db.DB(), "postgres", migrations, migrate.Up); err != nil {
		log.Fatalf("migrate err: %s", err)
	}
}

func GetDB() *gorm.DB {
	return db
}
