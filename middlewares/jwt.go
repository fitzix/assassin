package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/fitzix/assassin/models"
	"github.com/fitzix/assassin/service"
	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"
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

func JWT() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token, err := service.ParseToken(strings.TrimPrefix(c.Request().Header.Get("Authorization"), "Bearer "))
			if err != nil {
				return c.JSON(200, models.Response{
					Code: http.StatusUnauthorized,
					Msg:  http.StatusText(http.StatusUnauthorized),
					Data: nil,
				})
			}
			c.Set("token", token)
			return next(c)
		}
	}
}
