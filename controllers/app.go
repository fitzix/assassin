package controllers

import (
	"fmt"

	"github.com/fitzix/assassin/models"
	"github.com/fitzix/assassin/service"
	"github.com/gin-gonic/gin"
)

func AppGetAll(c *gin.Context) {
	a := service.NewAsnGin(c)
	var down models.AppList
	db := a.D
	if k := c.Query("key"); k != "" {
		db = db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", k))
	}

	if order := c.Query("order"); order == "hot" {
		db = db.Joins("LEFT JOIN app_hot ON app.id = app_hot.app_id").Order("app_hot.hot DESC")
	} else {
		db = db.Order("updated_at DESC")
	}

	if err := a.Page(db, &down.Apps, &down.Total); err != nil {
		a.Fail(service.StatusWebBadRequest)
		return
	}
	a.Success(down)
}

func AppIndex(c *gin.Context) {
	a := service.NewAsnGin(c)
	var down models.App
	if err := a.D.Find(&down, "id = ?", c.Param("id")).Error; err != nil {
		// a.log.Errorf("db err: %s", err)
		a.Fail(service.StatusWebBadRequest)
		return
	}
	a.Success(down)
}

func AppCreate(c *gin.Context) {}
func AppUpdate(c *gin.Context) {}
