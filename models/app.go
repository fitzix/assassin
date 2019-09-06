package models

import (
	"time"
)

type AsnModel struct {
	ID        string `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"-"`
}

type App struct {
	AsnModel
	Name        string `json:"name"`
	Icon        string `json:"icon"`
	Description string `json:"description"`
	Hot         int    `json:"hot"`
	View        int    `json:"view"`
}

type AppList struct {
	Total int   `json:"total"`
	Apps  []App `json:"apps"`
}

type CreateApp struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Tags        []int  `json:"tags"`
	Category    []int  `json:"category"`
}

type CreateAppVersion struct {
	AppID      string `json:"appId"`
	Version    string `json:"version"`
	Size       string `json:"size"`
	DownloadID int    `json:"downloadId"`
	URL        string `json:"url"`
	Secret     string `json:"secret"`
}
