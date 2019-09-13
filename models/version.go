package models

import "time"

type AppVersion struct {
	ID         int
	Name       string
	AppId      string
	Size       string
	DownloadId int
	Url        string
	Secret     string
	CreatedAt  time.Time
}

type Download struct {
	ID   uint
	Name string
}
