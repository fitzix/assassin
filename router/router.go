package router

import (
	"github.com/fitzix/assassin/controllers"
	"github.com/fitzix/assassin/middlewares"
	"github.com/gin-gonic/gin"
)

func Route(r *gin.Engine) {
	v1 := r.Group("/api/v1")

	auth := v1.Group("").Use(middlewares.JWT())
	{
		// app
		auth.POST("/apps", controllers.AppCreate)
		auth.PUT("/apps/:id", controllers.AppUpdate)
		auth.GET("/auth/apps", controllers.AppAuthorizedList)
		auth.GET("/auth/apps/:id", controllers.AppAuthorizedIndex)

		auth.PUT("/apps/:id/tags", controllers.AppTagsCreateOrUpdate)
		// version
		auth.POST("/apps/:id/versions", controllers.VersionCreate)
		auth.PUT("/apps/:id/versions/:versionId", controllers.VersionUpdate)
		// carousel
		auth.POST("/apps/:id/carousels", controllers.CarouselCreate)
		auth.DELETE("/apps/:id/carousels/:cid", controllers.CarouselDelete)

		// tag
		auth.POST("/tags", controllers.TagCreate)
		auth.PUT("/tags/:id", controllers.TagUpdate)
		// upload
		auth.POST("/upload", controllers.Upload)
	}

	public := v1.Group("")
	{
		// user
		public.POST("/login", controllers.UserLogin)
		auth.POST("/users", controllers.UserCreate)

		// app
		public.GET("/apps", controllers.AppList)
		public.GET("/apps/:id", controllers.AppIndex)

		// tag
		public.GET("/tags", controllers.TagList)
		// categories
		public.GET("/categories", controllers.CategoryList)
		// download types
		public.GET("/downloads", controllers.DownloadList)
	}
}
