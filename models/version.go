package models

import "time"

type AppVersion struct {
	ID int `gorm:"primary_key"`
	Name string
	Size string
	DownloadId int
	Url string
	Secret string
	CreatedAt time.Time
}