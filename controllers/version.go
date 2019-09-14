package controllers

import (
	"time"

	"github.com/fitzix/assassin/models"
	"github.com/fitzix/assassin/service"
	"github.com/gin-gonic/gin"
)

func AppVersion(c *gin.Context) {
	a := service.NewAsnGin(c)
	var down []models.AppVersion

	if err := a.D.Where("app_id = ?", c.Param("id")).Find(&down).Error; err != nil {
		a.Fail(service.StatusWebBadRequest, err)
		return
	}

	a.Success(down)
}

func VersionCreate(c *gin.Context) {
	a := service.NewAsnGin(c)
	var up models.AppVersion
	up.AppId = c.Param("id")

	if err := c.ShouldBind(&up); err != nil {
		a.Fail(service.StatusWebParamErr, err)
		return
	}
	tx := a.D.Begin()
	if err := tx.Create(&up).Error; err != nil {
		tx.Rollback()
		a.Fail(service.StatusWebBadRequest, err)
		return
	}
	if err := tx.Model(&models.App{}).Where("id = ?", up.AppId).Update("version_at", time.Now()).Error; err != nil {
		tx.Rollback()
		a.Fail(service.StatusWebBadRequest, err)
		return
	}
	tx.Commit()

	a.Success(up)
}

func VersionUpdate(c *gin.Context) {

}
