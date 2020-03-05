package controller

import (
	"github.com/fitzix/assassin/service"
	"github.com/labstack/echo/v4"
)

func Upload(e echo.Context) error {
	c := e.(*service.Context)
	file, err := c.FormFile("file")
	if err != nil {
		return c.Err(service.StatusWebParamErr, err)
	}
	if url, err := service.PutImage(file); err != nil {
		return c.Err(service.StatusWebBadRequest, err)
	} else {
		return c.Success(url)
	}
}
