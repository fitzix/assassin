package router

import (
	"github.com/fitzix/assassin/controllers"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	v1.GET("/", controllers.PlaygroundHandler())
	v1.POST("/query", controllers.GraphqlHandler())
}
