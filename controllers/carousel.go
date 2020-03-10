package controllers

import (
	"github.com/fitzix/assassin/models"
	"github.com/fitzix/assassin/service"
	"github.com/gin-gonic/gin"
	"github.com/t-tiger/gorm-bulk-insert"
)

func CarouseList(c *gin.Context) {
	a := service.NewAsnGin(c)
	var down []models.Carousel
	if err := a.D.Where("app_id = ?", c.Param("id")).Find(&down).Error; err != nil {
		a.Fail(service.StatusBadRequest, err)
		return
	}
	a.Success(down)
}

func CarouselCreate(c *gin.Context) {
	a := service.NewAsnGin(c)
	var up []string
	var insertRecords []interface{}
	if err := c.BindJSON(&up); err != nil {
		a.Fail(service.StatusParamErr, err)
		return
	}

	for _, v := range up {
		insertRecords = append(insertRecords, models.Carousel{
			AppId: c.Param("id"),
			Url:   v,
		})
	}

	if err := gormbulk.BulkInsert(a.D, insertRecords, 3000); err != nil {
		a.Fail(service.StatusBadRequest, err)
		return
	}

	a.Success(nil)
}

func CarouselDelete(c *gin.Context) {
	a := service.NewAsnGin(c)
	if err := a.D.Delete(&models.Carousel{}, "id = ?", c.Param("cid")).Error; err != nil {
		a.Fail(service.StatusBadRequest, err)
		return
	}
	a.Success(nil)
}

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
