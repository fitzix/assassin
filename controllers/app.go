package controllers

import (
	"strconv"

	"github.com/fitzix/assassin/models"
	"github.com/fitzix/assassin/service"
	"github.com/fitzix/assassin/utils"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/t-tiger/gorm-bulk-insert"
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
		db = db.Where("app.name LIKE %?%", req.Name)
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

	if service.AsnTypeExist(req.Order) {
		db = db.Joins("LEFT JOIN hot ON app.id = hot.app_id").Order("hot.hot DESC")
	} else {
		db = db.Order("app.updated_at DESC")
	}

	// rsp
	if err := db.Select("app.*, hot.hot, hot.view").Preload("Categories").Find(&rsp.Info).Error; err != nil {
		a.Fail(service.StatusBadRequest, err)
		return
	}

	a.SuccessWithPage(rsp)
}

func AppIndex(c *gin.Context) {
	a := service.NewAsnGin(c)
	var req models.AppIndexReq
	if err := c.ShouldBindUri(&req); err != nil {
		a.Fail(service.StatusParamErr, nil)
		return
	}
}

func AppAuthorizedIndex(c *gin.Context) {
	a := service.NewAsnGin(c)
	var down service.App

	db := a.D.Select("app.*, app_hot.*").Preload("Versions", func(db *gorm.DB) *gorm.DB {
		return db.Order("created_at DESC")
	}).Preload("Versions.AppVersionDownloads").Preload("Carousels").Preload("Tags").Joins("LEFT JOIN app_hot ON app.id = app_hot.app_id")

	if err := db.Find(&down, "app.id = ?", c.Param("id")).Error;
		err != nil {
		a.Fail(service.StatusBadRequest, nil)
		return
	}
	a.Success(down)
}

func AppCreate(c *gin.Context) {
	a := service.NewAsnGin(c)
	var up service.App
	if err := c.BindJSON(&up); err != nil {
		a.Fail(service.StatusParamErr, err)
		return
	}

	up.Versions = nil
	if len(up.Carousels) > 0 {
		up.Icon = up.Carousels[0].Url
		up.Carousels = up.Carousels[1:]
	}

	up.ID = utils.GenNanoId()

	if err := a.D.Omit("View", "Hot").Create(&up).Error; err != nil {
		a.Fail(service.StatusBadRequest, err)
		return
	}

	// create desc file to github
	// go func(id string) {
	// 	_, _ = service.GetGithubClient().CreateMdFile(id, service.AsnUploadTypeApp)
	// }(up.ID)

	a.Success(up)
}
func AppUpdate(c *gin.Context) {
	a := service.NewAsnGin(c)
	var up service.App

	if err := c.BindJSON(&up); err != nil {
		a.Fail(service.StatusParamErr, err)
		return
	}
	up.ID = c.Param("id")
	up.Versions = nil
	if len(up.Carousels) > 0 {
		up.Icon = up.Carousels[0].Url
		up.Carousels = up.Carousels[1:]
	}

	if err := a.D.Model(&up).Select("", up.CouldUpdateColumns()...).Updates(up).Error; err != nil {
		a.Fail(service.StatusBadRequest, err)
		return
	}
	a.Success(up)
}

func AppTags(c *gin.Context) {
	a := service.NewAsnGin(c)
	var down []models.AppTag
	if err := a.D.Where("app_id = ?", c.Param("id")).Find(&down).Error; err != nil {
		a.Fail(service.StatusBadRequest, err)
		return
	}
	a.Success(down)
}

func AppTagsCreateOrUpdate(c *gin.Context) {
	a := service.NewAsnGin(c)
	var up []uint
	if err := c.BindJSON(&up); err != nil {
		a.Fail(service.StatusParamErr, err)
		return
	}

	tx := a.D.Begin()

	if err := tx.Delete(&models.AppTag{}, "app_id IN ( ? )", up).Error; err != nil {
		tx.Rollback()
		a.Fail(service.StatusBadRequest, err)
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
		a.Fail(service.StatusBadRequest, err)
		return
	}
	tx.Commit()
	a.Success(insertRecords)
}
