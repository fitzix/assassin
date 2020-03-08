package middlewares

import (
	"net/http"
	"strings"

	"github.com/fitzix/assassin/models"
	"github.com/fitzix/assassin/service"
	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := service.ParseToken(strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer "))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusOK, models.Response{
				Code: http.StatusUnauthorized,
				Msg:  http.StatusText(http.StatusUnauthorized),
				Data: nil,
			})
			return
		}
		c.Set("token", token)
	}
}
