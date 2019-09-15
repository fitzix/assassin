package main

import (
	"log"

	"github.com/fitzix/assassin/middlewares"
	"github.com/fitzix/assassin/router"
	"github.com/fitzix/assassin/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	service.InitProject()
}

func main() {
	r := gin.New()

	r.Use(cors.Default())
	r.Use(middlewares.Zap(service.GetLogger()))
	r.Use(middlewares.ZapRecovery(service.GetLogger(), true))

	router.InitRouter(r)

	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}
