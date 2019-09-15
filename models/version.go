package models

import (
	"time"
)

type AppVersion struct {
	ID                  int                  `json:"id"`
	Name                string               `json:"name" binding:"required"`
	AppId               string               `json:"appId"`
	Size                string               `json:"size"`
	Status              bool                 `json:"status"`
	CreatedAt           time.Time            `json:"createdAt"`
	AppVersionDownloads []AppVersionDownload `json:"downloads"`
}

type Download struct {
	ID   uint
	Name string
}

type AppVersionDownload struct {
	ID           int    `json:"id"`
	AppVersionId int    `json:"appVersionId"`
	DownloadId   int    `json:"downloadId"`
	Url          string `json:"url"`
	Secret       string `json:"secret"`
}
