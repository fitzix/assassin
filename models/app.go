package models

import (
	"time"

	"github.com/fitzix/assassin/utils"
)

type AsnModel struct {
	ID        uint64     `json:"id"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"-"`
}

type App struct {
	AsnModel
	Name        string        `json:"name" binding:"required"`
	Type        int           `json:"type" binding:"oneof=0 1"`
	Icon        string        `json:"icon"`
	Title       string        `json:"title" binding:"required"`
	Description string        `json:"description"`
	Status      bool          `json:"status"`
	Hot         Hot           `json:"hot"`
	Categories  []AppCategory `json:"categories"`
	Tags        []AppTag      `json:"tags"`
	Versions    []Version     `json:"versions"`
	Carousels   []Carousel    `json:"carousels"`
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
	Info []App `json:"info"`
}

type AppIndexReq struct {
	Id uint64 `uri:"id" binding:"required"`
}

type AppVersionRsp struct {
	Version
	Sources []Source `json:"sources"`
}

func (a *App) Init() {
	a.Hot.Hot = 1
	a.Hot.View = 1
	a.UpdatedAt = time.Now()
	a.ID = utils.NextID()
}
