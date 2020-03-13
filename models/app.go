package models

import (
	"time"

	"github.com/fitzix/assassin/utils"
)

type AsnModel struct {
}

type App struct {
	ID          uint64     `json:"id"`
	CreatedAt   time.Time  `json:"-"`
	VersionAt   time.Time  `json:"versionAt"`
	DeletedAt   *time.Time `json:"-"`
	Name        string     `json:"name" binding:"required"`
	Type        int        `json:"type" binding:"oneof=0 1"`
	Icon        string     `json:"icon"`
	Title       string     `json:"title" binding:"required"`
	Description string     `json:"description"`
	Status      bool       `json:"status"`
}

type AppEdges struct {
	App
	Hot        Hot           `json:"hot" gorm:"foreignkey:app_id"`
	Categories []AppCategory `json:"categories" gorm:"foreignkey:app_id"`
	Tags       []AppTag      `json:"tags" gorm:"foreignkey:app_id"`
	Versions   []Version     `json:"versions" gorm:"foreignkey:app_id"`
	Carousels  []Carousel    `json:"carousels" gorm:"foreignkey:app_id"`
}

// gorm:"many2many:app_category;jointable_foreignkey:app_id;"

// func (a *App) CouldUpdateColumns() []interface{} {
// 	return []interface{}{
// 		"name",
// 		"type",
// 		"icon",
// 		"description",
// 		"status",
// 		"update_at",
// 	}
// }

type AppListReq struct {
	Name   string `json:"name" form:"name"`
	Order  string `json:"order" form:"order" binding:"omitempty,eq=hot"`
	Type   string `json:"type" form:"type" binding:"omitempty,oneof=app book"`
	Status string `json:"status" form:"status" binding:"omitempty,oneof=pub unpub"`
	PageReq
}

type AppListRsp struct {
	PageRsp
	Info []AppEdges `json:"info"`
}

type AppIndexReq struct {
	Id uint64 `uri:"id" binding:"required"`
}

type AppVersionRsp struct {
	Version
	Sources []Source `json:"sources"`
}

// table name
func (AppEdges) TableName() string {
	return "app"
}

// hooks
func (a *App) BeforeCreate() {
	a.ID = utils.NextID()
	a.VersionAt = time.Now()
}
