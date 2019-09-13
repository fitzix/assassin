package main

import (
	"log"

	"github.com/fitzix/assassin/db"
	"github.com/fitzix/assassin/middlewares"
	"github.com/fitzix/assassin/router"
	"github.com/fitzix/assassin/utils"
	"github.com/fitzix/assassin/utils/github"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	utils.InitConf()
	db.Init()
	utils.InitLogger()
	github.InitGithubClient()
}

func main() {
	r := gin.New()

	r.Use(cors.Default())
	r.Use(middlewares.Zap(utils.GetLogger()))
	r.Use(middlewares.ZapRecovery(utils.GetLogger(), true))

	router.InitRouter(r)

	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}
