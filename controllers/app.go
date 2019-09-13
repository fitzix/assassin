package controllers

import (
	"fmt"

	"github.com/fitzix/assassin/models"
	"github.com/fitzix/assassin/service"
	"github.com/gin-gonic/gin"
	gormbulk "github.com/t-tiger/gorm-bulk-insert"
)

func AppList(c *gin.Context) {
	a := service.NewAsnGin(c)
	var down []models.App
	var total int

	db := a.D.Select("app.*, app_hot.*").Joins("LEFT JOIN app_hot ON app.id = app_hot.app_id").Where("app.status", true)
	if k := c.Query("key"); k != "" {
		db = db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", k))
	}

	if order := c.Query("order"); order == "hot" {
		db = db.Order("app_hot.hot DESC")
	} else {
		db = db.Order("app.updated_at DESC")
	}

	if appType := c.Query("type"); appType == "app" {
		db = db.Where("type = ?", 1)
	} else {
		db = db.Where("type = ?", 0)
	}

	if err := a.Page(db, &down, &total); err != nil {
		a.Fail(service.StatusWebBadRequest, nil)
		return
	}
	a.SuccessWithPage(total, down)
}

func AppIndex(c *gin.Context) {
	a := service.NewAsnGin(c)
	var down models.App

	if err := a.D.Select("app.*, app_hot.*").Joins("LEFT JOIN app_hot ON app.id = app_hot.app_id").Where("app.status", true).Find(&down, "app.id = ?", c.Param("id")).Error; err != nil {
		a.L.Errorf("err %s", err)
		a.Fail(service.StatusWebBadRequest, nil)
		return
	}
	a.Success(down)
}

func AppCreate(c *gin.Context) {
	a := service.NewAsnGin(c)
	var up models.App
	if err := c.BindJSON(&up); err != nil {
		a.Fail(service.StatusWebParamErr, err)
		return
	}

	if err := a.D.Omit("View", "Hot").Create(&up).Error; err != nil {
		a.Fail(service.StatusWebBadRequest, err)
		return
	}

	a.Success(up)
}
func AppUpdate(c *gin.Context) {
	a := service.NewAsnGin(c)
	var up models.App
	up.ID = c.Param("id")

	if err := c.BindJSON(&up); err != nil {
		a.Fail(service.StatusWebParamErr, err)
		return
	}
	if err :=a.D.Model(&up).Select("", up.CouldUpdateColumns()...).Updates(up).Error; err != nil {
		a.Fail(service.StatusWebBadRequest, err)
		return
	}
	a.Success(up)
}

func AppTags(c *gin.Context) {
	a := service.NewAsnGin(c)
	var down []models.AppTag
	if err := a.D.Where("app_id = ?", c.Param("id")).Find(&down).Error; err != nil {
		a.Fail(service.StatusWebBadRequest, err)
		return
	}
	a.Success(down)
}

func AppTagsCreateOrUpdate(c *gin.Context) {
	a := service.NewAsnGin(c)
	var up []uint
	if err := c.BindJSON(&up); err != nil {
		a.Fail(service.StatusWebParamErr, err)
		return
	}

	tx := a.D.Begin()

	if err := tx.Delete(&models.AppTag{}, "app_id IN ( ? )", up).Error; err != nil {
		tx.Rollback()
		a.Fail(service.StatusWebBadRequest, err)
		return
	}

	var insertRecords []interface{}

	for _, v := range up {
		insertRecords = append(insertRecords, models.AppTag{
			AppId: c.Param("id"),
			TagId: v,
		})
	}

	if err := gormbulk.BulkInsert(tx, insertRecords, 3000); err != nil {
		tx.Rollback()
		a.Fail(service.StatusWebBadRequest, err)
		return
	}
	tx.Commit()
	a.Success(insertRecords)
}