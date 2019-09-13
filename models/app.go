package models

import (
	"time"

	"github.com/fitzix/assassin/utils/encrypt"
	"github.com/jinzhu/gorm"
)

type AsnModel struct {
	ID        string     `json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `sql:"index" json:"-"`
}

type App struct {
	ID          string        `json:"id"`
	Name        string        `json:"name"`
	Type        int           `json:"type"`
	Icon        string        `json:"icon"`
	Description string        `json:"description"`
	Status      bool          `json:"status"`
	View        int           `json:"view"`
	Hot         int           `json:"hot"`
	Carousels   []AppCarousel `json:"carousels"`
	Versions    []AppVersion  `json:"versions"`
	CreatedAt   time.Time     `json:"-"`
	UpdatedAt   time.Time     `json:"-"`
	DeletedAt   *time.Time    `json:"-"`
}

type AppCarousel struct {
	ID    uint   `json:"id"`
	AppId string `json:"appId"`
	Url   string `json:"url"`
}

type CreateAppVersion struct {
	AppID      string
	Version    string
	Size       string
	DownloadID int
	URL        string
	Secret     string
}

// hook
func (a *App) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("ID", encrypt.GetNanoId())
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
