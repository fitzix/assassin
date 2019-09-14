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
	ID    uint
	AppId string
	TagId uint
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
