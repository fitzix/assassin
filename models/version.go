package models

import "time"

type Version struct {
	ID int `gorm:"primary_key"`
	Name string
	DownloadId int
	Url string
	Secret string
	CreatedAt time.Time
}