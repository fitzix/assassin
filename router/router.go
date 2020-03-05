package router

import (
	"github.com/fitzix/assassin/controller"
	"github.com/fitzix/assassin/controllers"
	"github.com/fitzix/assassin/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"
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
	authGroup.GET("/auth/apps", controllers.AppAuthorizedList)
	authGroup.GET("/auth/apps/:id", controllers.AppAuthorizedIndex)

	authGroup.PUT("/apps/:id/tags", controllers.AppTagsCreateOrUpdate)
	// version
	authGroup.POST("/apps/:id/versions", controllers.VersionCreate)
	authGroup.PUT("/apps/:id/versions/:versionId", controllers.VersionUpdate)
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
	publicGroup.POST("/login", controllers.UserLogin)

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

func Route(e *echo.Echo) {
	v1 := e.Group("/api/v1")

	auth := v1.Group("")
	auth.Use(middlewares.JWT())
	{
		// upload
		auth.POST("/uploads/img", controller.Upload)
	}

	public := v1.Group("")
	{
		public.POST("/login", controller.UserLogin)
		public.POST("/users", controller.UserCreate)
		public.GET("/apps", controller.AppList)
	}
}
