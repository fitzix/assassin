package service

import "github.com/fitzix/assassin/models"

type Token struct {
	Uid  string `json:"uid"`
	Code uint   `json:"code"`
}

type UserLogin struct {
	UserName string `json:"name" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type App struct {
	models.App
	View      int                  `json:"view"`
	Hot       int                  `json:"hot"`
	Versions  []AppVersion         `json:"versions"`
	Tags      []models.AppTag      `json:"tags"`
	Carousels []models.AppCarousel `json:"carousels"`
}

type AppVersion struct {
	models.AppVersion
	AppVersionDownloads []AppVersionDownload `json:"downloads"`
}

type AppVersionDownload struct {
	models.AppVersionDownload
}

func (v *AppVersionDownload) AfterFind() (err error) {
	if v.Secret != "" {
		v.Secret = AesEncrypt(v.Secret)
	}
	return
}
