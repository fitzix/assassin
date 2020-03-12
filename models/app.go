package models

import (
	"time"
)

type AsnModel struct {
	ID        string     `json:"id"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"-"`
}

type App struct {
	AsnModel
	Name   string `json:"name"`
	Type   int    `json:"type"`
	Icon   string `json:"icon"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

// gorm:"many2many:app_category;jointable_foreignkey:app_id;"

type AppList struct {
	App
	Hot        int           `json:"hot"`
	View       int           `json:"view"`
	Categories []AppCategory `json:"categories" gorm:"foreignkey:app_id"`
}

type AppTag struct {
	ID    uint   `json:"id"`
	AppId string `json:"appId"`
	TagId uint   `json:"tagId"`
}

func (a *App) CouldUpdateColumns() []interface{} {
	return []interface{}{
		"name",
		"type",
		"icon",
		"description",
		"status",
		"update_at",
	}
}

type AppListReq struct {
	Name   string `json:"name" form:"name"`
	Order  string `json:"order" form:"order" binding:"omitempty,eq=hot"`
	Type   string `json:"type" form:"type" binding:"omitempty,oneof=app book"`
	Status string `json:"status" form:"status" binding:"omitempty,oneof=pub unpub"`
	PageReq
}

type AppListRsp struct {
	PageRsp
	Info []AppList `json:"info"`
}

type AppIndexReq struct {
	Id int `uri:"id" binding:"required"`
}

type AppVersionRsp struct {
	Version
	Sources []Source `json:"sources"`
}
