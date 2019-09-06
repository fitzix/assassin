package router

import (
	"github.com/fitzix/assassin/controllers"
	"github.com/fitzix/assassin/middlewares"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	v1 := r.Group("/api/v1")

	authGroup := v1.Group("")
	authGroup.Use(middlewares.Jwt(""))

	publicGroup := v1.Group("")

	// ===================================

	authGroup.POST("/apps", controllers.AppCreate)
	authGroup.PUT("/apps/:id", controllers.AppUpdate)

	authGroup.POST("/tags", controllers.TagCreate)
	authGroup.PUT("/tags/:id", controllers.TagUpdate)

	// ===================================

	publicGroup.GET("/apps", controllers.AppGetAll)
	publicGroup.GET("/apps/:id", controllers.AppIndex)

	publicGroup.GET("/tags", controllers.TagAll)
	publicGroup.GET("/tags/:id", controllers.TagIndex)

	publicGroup.GET("/versions", controllers.TagAll)

	publicGroup.GET("/categories", controllers.CategoryAll)
}
