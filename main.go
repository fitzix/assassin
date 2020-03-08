package main

import (
	"github.com/fitzix/assassin/middlewares"
	"github.com/fitzix/assassin/router"
	"github.com/fitzix/assassin/service"
	"github.com/gin-gonic/gin"
)

func main() {
	service.Init()

	r := gin.New()
	r.Use(middlewares.ZapRecovery(service.GetLogger(), true))
	r.Use(middlewares.Zap(service.GetLogger()))
	r.Use(middlewares.CORS())

	router.Route(r)

	if err := r.Run(); err != nil {
		service.GetLogger().Fatal(err)
	}
}
