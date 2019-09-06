package router

import (
	"github.com/fitzix/assassin/controllers"
	"github.com/gin-gonic/gin"
)

func appRoute(g *gin.RouterGroup) {
	g.GET("", controllers.AppGetAll)
	g.GET("/:id")
	g.POST("", controllers.AppCreate)
}
