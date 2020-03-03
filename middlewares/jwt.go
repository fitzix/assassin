package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/fitzix/assassin/models"
	"github.com/fitzix/assassin/service"
	"github.com/gin-gonic/gin"
)

// TODO check token code
func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := service.ParseToken(strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer "))
		if err != nil {
			fmt.Printf("%s", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, models.Response{
				Code: http.StatusUnauthorized,
				Msg:  http.StatusText(http.StatusUnauthorized),
				Data: nil,
			})
			return
		}
		c.Set("token", token)
	}
}
