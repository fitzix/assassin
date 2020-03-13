package controllers

import (
	"fmt"
	"strconv"

	"github.com/fitzix/assassin/models"
	"github.com/fitzix/assassin/service"
	"github.com/gin-gonic/gin"
	gormbulk "github.com/t-tiger/gorm-bulk-insert"
)

func AppList(c *gin.Context) {
	a := service.NewAsnGin(c)
	var req models.AppListReq
	if err := c.ShouldBindQuery(&req); err != nil {
		a.Fail(service.StatusParamErr, err)
		return
	}

	// 查询未发布
	if service.AsnType(req.Status) < 1 {
		// 未授权 只能查询已发布
		if !a.IsAuth() {
			req.Status = "pub"
		}
	}

	var rsp models.AppListRsp

	db := a.D.Table("app")

	if req.Name != "" {
		db = db.Where("app.name LIKE ?", fmt.Sprintf("%%%s%%", req.Name))
	}

	if t := service.AsnType(req.Type); t > -1 {
		db = db.Where("app.type = ?", t)
	}

	if t := service.AsnType(req.Status); t > -1 {
		db = db.Where("app.status = ?", strconv.Itoa(t))
	}

	// count
	if err := db.Table("app").Count(&rsp.Total).Error; err != nil {
		a.Fail(service.StatusBadRequest, err)
		return
	}

	if rsp.Total == 0 {
		a.SuccessWithPage(rsp)
		return
	}

	if service.AsnTypeExist(req.Order) {
		db = db.Joins("LEFT JOIN hot ON app.id = hot.app_id").Order("hot.hot DESC")
	} else {
		db = db.Order("app.updated_at DESC")
	}

	// rsp
	if err := a.Page(db.Preload("Categories").Preload("Hot"), req.PageReq, &rsp.Info); err != nil {
		a.Fail(service.StatusBadRequest, err)
		return
	}

	a.SuccessWithPage(rsp)
}

func AppIndex(c *gin.Context) {
	a := service.NewAsnGin(c)
	var req models.AppIndexReq
	if err := c.ShouldBindUri(&req); err != nil {
		a.Fail(service.StatusParamErr, err)
		return
	}
	var rsp models.AppEdges
	db := a.D.
		Where("id = ?", req.Id).
		Preload("Hot").
		Preload("Carousels").
		Preload("Categories").
		Preload("Tags").
		Preload("Versions.Sources")
	if !a.IsAuth() {
		db = db.Where("app.status = ?", true)
	}
	if err := db.First(&rsp).Error; err != nil {
		a.Fail(service.StatusBadRequest, err)
		return
	}
	a.Success(rsp)
}

func AppCreate(c *gin.Context) {
	a := service.NewAsnGin(c)
	var app models.AppEdges
	if err := c.ShouldBindJSON(&app); err != nil {
		a.Fail(service.StatusParamErr, err)
		return
	}
	if err := a.D.Create(&app).Error; err != nil {
		a.Fail(service.StatusBadRequest, err)
		return
	}

	a.Success(app)
}
func AppUpdate(c *gin.Context) {
	a := service.NewAsnGin(c)
	var uri models.AppIndexReq
	var req models.App
	if c.ShouldBindJSON(&req) != nil || c.ShouldBindUri(&uri) != nil {
		a.Fail(service.StatusParamErr, nil)
		return
	}
	req.ID = uri.Id
	if err := a.D.Unscoped().Model(&req).Omit("id", "version_at").Updates(req).Error; err != nil {
		a.Fail(service.StatusBadRequest, err)
		return
	}
	a.Success(req)
}

func AppTagsCreateOrUpdate(c *gin.Context) {
	a := service.NewAsnGin(c)
	var req []uint
	var uri models.AppIndexReq
	if c.ShouldBindJSON(&req) != nil || c.ShouldBindUri(&uri) != nil {
		a.Fail(service.StatusParamErr, nil)
		return
	}
	tx := a.D.Begin()
	if err := tx.Delete(&models.AppTag{}, "app_id = ?", uri.Id).Error; err != nil {
		tx.Rollback()
		a.Fail(service.StatusBadRequest, err)
		return
	}
	var appTags []interface{}
	for _, v := range req {
		appTags = append(appTags, models.AppTag{
			AppID: uri.Id,
			TagID: v,
		})
	}
	if err := gormbulk.BulkInsert(tx, appTags, len(appTags)); err != nil {
		tx.Rollback()
		a.Fail(service.StatusBadRequest, err)
		return
	}
	tx.Commit()
	a.Success(appTags)
}
