package controllers

import (
	"github.com/fitzix/assassin/service"
	"github.com/gin-gonic/gin"
)

func Upload(c *gin.Context) {
	a := service.NewAsnGin(c)
	file, err := c.FormFile("file")
	if err != nil {
		a.Fail(service.StatusParamErr, err)
		return
	}
	url, err := service.PutImage(file)
	if err != nil {
		a.Fail(service.StatusBadRequest, err)
		return
	}
	a.Success(url)
}
