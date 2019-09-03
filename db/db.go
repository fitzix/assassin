package db

import (
	"fmt"
	"log"

	"github.com/fitzix/assassin/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rubenv/sql-migrate"
)

var db *sqlx.DB

func Init(conf models.Config) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", conf.Db.Host, conf.Db.Port, conf.Db.User, conf.Db.Password, conf.Db.Dbname)
	db = sqlx.MustConnect("postgres", connStr)
	migrateDb()
}

func migrateDb() {
	migrations := &migrate.FileMigrationSource{Dir: "db/migrations"}
	if _, err := migrate.Exec(db.DB, "postgres", migrations, migrate.Up); err != nil {
		log.Fatalf("migrate err: %s", err)
	}
}

func GetDB() *sqlx.DB {
	return db
}
