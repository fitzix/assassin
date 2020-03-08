package controllers

import (
	"github.com/fitzix/assassin/models"
	"github.com/fitzix/assassin/service"
	"github.com/gin-gonic/gin"
)

func TagList(c *gin.Context) {
	a := service.NewAsnGin(c)
	var down []models.Tag
	if err := a.D.Find(&down).Error; err != nil {
		a.Fail(service.StatusBadRequest, err)
		return
	}
	a.Success(down)
}
func TagIndex(c *gin.Context) {

}
func TagCreate(c *gin.Context) {
	a := service.NewAsnGin(c)
	var up models.Tag
	if err := c.ShouldBindJSON(&up); err != nil {
		a.Fail(service.StatusParamErr, err)
		return
	}

	if err := a.D.Create(&up).Error; err != nil {
		a.Fail(service.StatusBadRequest, err)
		return
	}

	a.Success(up)
}
func TagUpdate(c *gin.Context) {}
func TagDelete(c *gin.Context) {}
