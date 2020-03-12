package models

import (
	"time"

	"github.com/fitzix/assassin/ent"
)

type AsnModel struct {
	ID        uint       `json:"-" db:"id"`
	CreatedAt time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time  `json:"updatedAt" db:"updated_at"`
	DeletedAt *time.Time `json:"-" db:"deleted_at"`
}

type App struct {
	AppID    string `json:"appId" db:"app_id"`
	Name     string `json:"name"`
	Type     int    `json:"type"`
	Icon     string `json:"icon" db:"icon"`
	Title    string `json:"title" db:"title"`
	Status   bool   `json:"status" db:"status"`
	Category int    `json:"category" db:"category"`
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
	Order  string `json:"order" form:"order"`
	Type   string `json:"type" form:"type" validate:"oneof=app book"`
	Status string `json:"status" form:"status" validate:"oneof=pub unpub"`
	PageReq
}

type AppListRsp struct {
	PageRsp
	Info ent.Apps `json:"info"`
}

type AppIndexReq struct {
	Id int `uri:"id" binding:"required"`
}

type AppVersionRsp struct {
	Version
	Sources []Source `json:"sources"`
}
