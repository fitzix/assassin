package router

import (
	"github.com/fitzix/assassin/controllers"
	"github.com/fitzix/assassin/middlewares"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	v1 := r.Group("/api/v1")

	// =================需要登录==================
	authGroup := v1.Group("")
	authGroup.Use(middlewares.Jwt())

	// user
	authGroup.POST("/users", controllers.UserCreate)

	// app
	authGroup.POST("/apps", controllers.AppCreate)
	authGroup.PUT("/apps/:id", controllers.AppUpdate)

	authGroup.PUT("/apps/:id/tags", controllers.AppTagsCreateOrUpdate)
	// version
	authGroup.POST("/apps/:id/version", controllers.VersionCreate)
	authGroup.PUT("/apps/:id/version/:versionId", controllers.VersionUpdate)
	// carousel
	authGroup.POST("/apps/:id/carousels", controllers.CarouselCreate)
	authGroup.DELETE("/apps/:id/carousels/:cid", controllers.CarouselDelete)

	// tag
	authGroup.POST("/tags", controllers.TagCreate)
	authGroup.PUT("/tags/:id", controllers.TagUpdate)
	// upload
	authGroup.POST("/upload", controllers.Upload)

	// ================公开接口===================
	publicGroup := v1.Group("")

	// user
	publicGroup.POST("/users/login", controllers.UserLogin)

	// app
	publicGroup.GET("/apps", controllers.AppList)
	publicGroup.GET("/apps/:id", controllers.AppIndex)

	// tag
	publicGroup.GET("/tags", controllers.TagList)
	// categories
	publicGroup.GET("/categories", controllers.CategoryList)
	// download types
	publicGroup.GET("/downloads", controllers.DownloadList)
}
