package controllers

import (
	"context"
	"fmt"

	"github.com/fitzix/assassin/ent"
	"github.com/fitzix/assassin/ent/app"
	"github.com/fitzix/assassin/ent/version"
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
		a.Fail(service.StatusParamErr, nil)
		return
	}

	// 查询未发布
	if service.AsnType(req.Status) < 1 {
		// 未授权 只能查询已发布
		if !a.IsAuth() {
			req.Status = "pub"
		}
	}

	if service.AsnType(req.Order) > 0 {
		appListHot(a, req)
		return
	}

	query := a.Db.App.Query()

	if req.Name != "" {
		query.Where(app.NameContains(fmt.Sprintf("%%%s%%", req.Name)))
	}

	if t := service.AsnType(req.Type); t > -1 {
		query.Where(app.TypeEQ(t))
	}

	if t := service.AsnType(req.Status); t > -1 {
		query.Where(app.TypeEQ(t))
	}
	ctx := context.Background()
	total, err := query.Count(ctx)
	if err != nil {
		a.Fail(service.StatusBadRequest, err)
		return
	}

	apps, err := query.Order(ent.Desc(app.FieldUpdatedAt)).WithHot().WithCategories().All(ctx)
	if err != nil {
		a.Fail(service.StatusBadRequest, err)
		return
	}
	a.SuccessWithPage(total, apps)
}

func appListHot(a *service.AsnGin, req models.AppListReq) {
	var rsp models.AppListRsp
	if err := service.GetSqlDB().QueryRow(`SELECT count(*) FROM "app" WHERE "app"."type" = $1 AND "app"."status" = $2`, service.AsnType(req.Type), service.AsnAppStatusPublish).Scan(&rsp.Total); err != nil {
		a.Fail(service.StatusBadRequest, err)
		return
	}
	service.GetSqlDB().Query(`SELECT * FROM "app" LEFT JOIN "hot" ON "app"."id" = "hot"."app_id" WHERE "app"."type" = $1 AND "app"."status" = $2 `, service.AsnType(req.Type), service.AsnAppStatusPublish)
}

func AppIndex(c *gin.Context) {
	a := service.NewAsnGin(c)
	var req models.AppIndexReq
	if err := c.ShouldBindUri(&req); err != nil {
		a.Fail(service.StatusParamErr, nil)
		return
	}
	query := a.Db.App.Query().Where(app.IDEQ(req.Id)).WithHot().WithCategories().WithTags()
	if a.IsAuth() {
		query.WithVersions(func(q *ent.VersionQuery) {
			q.WithSources(func(q *ent.SourceQuery) {
				q.WithProvider()
			})
		})
	} else {
		query.Where(app.StatusEQ(0))
		query.WithVersions(func(q *ent.VersionQuery) {
			q.Where(version.StatusEQ(0))
			q.WithSources(func(q *ent.SourceQuery) {
				q.WithProvider()
			})
		})
	}

	rsp, err := query.Only(context.Background())
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
