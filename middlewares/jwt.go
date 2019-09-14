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
		tokenSlice := strings.Split(c.GetHeader("Authorization"), " ")
		if len(tokenSlice) < 2 {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		token, err := encrypt.ParseToken(tokenSlice[1])
		if err != nil {
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
