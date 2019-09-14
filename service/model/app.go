package model

import "github.com/fitzix/assassin/models"

type App struct {
	models.App
	View      int                  `json:"view"`
	Hot       int                  `json:"hot"`
	Versions  []models.AppVersion  `json:"versions"`
	Tags      []models.AppTag      `json:"tags"`
	Carousels []models.AppCarousel `json:"carousels"`
}
