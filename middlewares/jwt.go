package middlewares

import (
	"net/http"
	"strings"

	"github.com/fitzix/assassin/models"
	"github.com/fitzix/assassin/utils/encrypt"
	"github.com/gin-gonic/gin"
)

// TODO check token code
func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := strings.Trim(c.GetHeader("Authorization"), "Bear ")
		token, err := encrypt.ParseToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, models.Response{
				Code: http.StatusUnauthorized,
				Msg:  http.StatusText(http.StatusUnauthorized),
				Data: nil,
			})
			c.Abort()
			return
		}
		c.Set("token", token)
	}
}
