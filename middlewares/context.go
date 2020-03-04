package middlewares

import (
	"github.com/fitzix/assassin/service"
	"github.com/labstack/echo/v4"
)

func Context() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return next(&service.Context{Context: c})
		}
	}
}
