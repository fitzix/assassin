package models

import (
	"time"
)

type AsnModel struct {
	ID        string     `json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `sql:"index" json:"-"`
}

type App struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Type      int        `json:"type"`
	Icon      string     `json:"icon"`
	Title     string     `json:"title"`
	Status    bool       `json:"status"`
	Category  int        `json:"category"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
	VersionAt time.Time  `json:"versionAt"`
}

type AppCarousel struct {
	ID    uint   `json:"id"`
	AppId string `json:"appId"`
	Url   string `json:"url"`
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

type PageReq struct {
	PageSize int `json:"pageSize" query:"pageSize"`
	PageNum  int `json:"pageNum" query:"pageNum"`
}

type AppListReq struct {
	Name  string `query:"name"`
	Order string `query:"order"`
	PageReq
}
