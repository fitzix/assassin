package main

import (
	"log"

	"github.com/fitzix/assassin/db"
	"github.com/fitzix/assassin/middlewares"
	"github.com/fitzix/assassin/models"
	"github.com/fitzix/assassin/router"
	"github.com/fitzix/assassin/utils"
	"github.com/gin-gonic/gin"
)

var AppConf models.Config

func main() {
	if err := utils.LoadConfig(&AppConf); err != nil {
		log.Fatalf("load config err: %s", err)
	}
	db.Init(AppConf)
	utils.InitLogger(gin.Mode())

	r := gin.New()
	r.Use(middlewares.Zap(utils.GetLogger()), middlewares.ZapRecovery(utils.GetLogger(), true))
	router.InitRouter(r)

	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}