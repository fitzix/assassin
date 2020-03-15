package service

import "github.com/fitzix/assassin/models"


type AppVersion struct {
	models.Version
	AppVersionDownloads []AppVersionDownload `json:"downloads"`
}

type AppVersionDownload struct {
	models.Source
}

func (v *AppVersionDownload) AfterFind() (err error) {
	// if v.Secret != "" {
	// 	v.Secret = AesEncrypt(v.Secret)
	// }
	return
}
