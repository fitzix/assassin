package main

import (
	"log"

	"github.com/fitzix/assassin/middlewares"
	"github.com/fitzix/assassin/router"
	"github.com/fitzix/assassin/service"
	"github.com/gin-gonic/gin"
)

func init() {
	service.InitProject()
}

func main() {
	r := gin.New()

	r.Use(middlewares.Cors())

	router.InitRouter(r)

	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}
