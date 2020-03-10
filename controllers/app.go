package controllers

import (
	"fmt"
	"strconv"

	"github.com/fitzix/assassin/models"
	"github.com/fitzix/assassin/schema"
	"github.com/fitzix/assassin/service"
	"github.com/fitzix/assassin/utils"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/t-tiger/gorm-bulk-insert"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

func appList(c *gin.Context, isAuth bool) {
	a := service.NewAsnGin(c)
	var req models.AppListReq
	if err := c.ShouldBindQuery(&req); err != nil {
		a.Fail(service.StatusParamErr, nil)
		return
	}

	if !isAuth {
		req.Status = "pub"
	}
	var query []qm.QueryMod
	if req.Name != "" {
		query = append(query, qm.Where("app.name LIKE ?", fmt.Sprintf("%%%s%%", req.Name)))
	}

	if t := service.AsnType(req.Type); t > -1 {
		query = append(query, qm.Where("app.type = ?", t))
	}

	if status := service.AsnType(req.Status); status > -1 {
		query = append(query, qm.Where("app.status = ?", strconv.Itoa(status)))
	}
	total, err := schema.Apps(query...).Count(a.Db)
	if err != nil {
		a.Fail(service.StatusBadRequest, err)
		return
	}

	if service.AsnTypeExist(req.Order) {
		query = append(query, qm.OrderBy("hot.hot DESC"), qm.LeftOuterJoin("hot ON app.hot_id = hot.id"))
	} else {
		qm.OrderBy("app.update_at DESC")
	}

	apps, err := schema.Apps(append(query, qm.Load(schema.AppRels.Hot), qm.Load(schema.AppRels.Tags))...).All(a.Db)
	if err != nil {
		a.Fail(service.StatusBadRequest, err)
		return
	}
	var rsp []models.AppCover
	for _, app := range apps {
		rsp = append(rsp, models.AppCover{
			App:  app,
			Hot:  app.R.Hot,
			Tags: app.R.Tags,
		})
	}

	a.SuccessWithPage(total, rsp)
}

func appIndex(c *gin.Context, isAuth bool) {
	a := service.NewAsnGin(c)
	var rsp models.AppIndex
	db := a.D.Select("app.*, app_hot.*")
	// if !isAuth {
	// 	db = db.Where()
	// }
	err := db.
		Preload("Versions", func(db *gorm.DB) *gorm.DB {
			if !isAuth {
				db = db.Where("status = ?", true)
			}
			return db.Order("created_at DESC")
		}).
		Preload("Versions.AppVersionDownloads").
		Preload("Carousels").
		Preload("Tags").
		Joins("LEFT JOIN app_hot ON app.id = app_hot.app_id").
		Where("app.status = ?", true).
		Find(&rsp, "app.id = ? AND ", c.Param("id")).Error
	if err != nil {
		a.Fail(service.StatusBadRequest, err)
		return
	}
	a.Success(rsp)
}

func AppList(c *gin.Context) {
	appList(c, false)
}

// AppAuthorizedList need auth
func AppAuthorizedList(c *gin.Context) {
	appList(c, true)
}

func AppIndex(c *gin.Context) {
	a := service.NewAsnGin(c)
	var rsp models.AppIndex
	err := a.D.
		Select("app.*, app_hot.*").
		Preload("Versions", func(db *gorm.DB) *gorm.DB {
			return db.Where("status = ?", true).Order("created_at DESC")
		}).
		Preload("Versions.AppVersionDownloads").
		Preload("Carousels").Preload("Tags").
		Joins("LEFT JOIN app_hot ON app.id = app_hot.app_id").
		Where("app.status = ?", true).
		Find(&rsp, "app.id = ?", c.Param("id")).Error
	if err != nil {
		a.Fail(service.StatusBadRequest, err)
		return
	}
	a.Success(rsp)
}

func AppAuthorizedIndex(c *gin.Context) {
	a := service.NewAsnGin(c)
	var down service.App

	db := a.D.Select("app.*, app_hot.*").Preload("Versions", func(db *gorm.DB) *gorm.DB {
		return db.Order("created_at DESC")
	}).Preload("Versions.AppVersionDownloads").Preload("Carousels").Preload("Tags").Joins("LEFT JOIN app_hot ON app.id = app_hot.app_id")

	if err := db.Find(&down, "app.id = ?", c.Param("id")).Error;
		err != nil {
		a.L.Errorf("err %s", err)
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

	up.AppID = utils.GenNanoId()

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
	up.AppID = c.Param("id")
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
