package models

import "time"

type AppVersion struct {
	ID         int       `json:"id"`
	Name       string    `json:"name" binding:"required"`
	AppId      string    `json:"appId"`
	Size       string    `json:"size"`
	DownloadId int       `json:"downloadId" binding:"required"`
	Url        string    `json:"url"`
	Secret     string    `json:"secret"`
	CreatedAt  time.Time `json:"createdAt"`
}

type Download struct {
	ID   uint
	Name string
}
