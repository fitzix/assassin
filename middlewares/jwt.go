package middlewares

import (
	"github.com/gin-gonic/gin"
)

func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.GetHeader("Bearer ")
	}
}