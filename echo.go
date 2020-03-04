package main

import (
	"context"
	"fmt"
	"log"

	"github.com/fitzix/assassin/ent"
	"github.com/fitzix/assassin/ent/migrate"
	_ "github.com/lib/pq"
)

func main() {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "127.0.0.1", 5432, "fitz", "131833", "assassin-ent")
	client, err := ent.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	// run the auto migration tool.
	if err := client.Schema.Create(
		context.Background(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	// e := echo.New()
	// // Middleware
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())
	//
	// // Routes
	// e.GET("/", hello)
	//
	// // Start server
	// e.Logger.Fatal(e.Start(":1323"))
}
