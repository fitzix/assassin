package model

import "github.com/fitzix/assassin/models"

type App struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Type        int    `json:"type"`
	Icon        string `json:"icon"`
	Description string `json:"description"`
	View        int    `json:"view"`
	Hot         int    `json:"hot"`
}

type IndexApp struct {
	App
	Carousels []models.AppCarousel `json:"carousels" gorm:"foreignkey:AppID;"`
	Versions  []models.AppVersion  `json:"versions" gorm:"foreignkey:AppID;"`
}

func (IndexApp) TableName() string {
	return "app"
}
