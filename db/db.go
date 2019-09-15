package db

import (
	"database/sql"
	"log"

	"github.com/rubenv/sql-migrate"
)


func MigrateDb(db *sql.DB) {
	migrations := &migrate.FileMigrationSource{Dir: "db/migrations"}
	if _, err := migrate.Exec(db, "postgres", migrations, migrate.Up); err != nil {
		log.Fatalf("migrate err: %s", err)
	}
}
