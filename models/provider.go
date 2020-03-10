package models

type Provider struct {
	ID           int    `json:"id"`
	AppVersionId int    `json:"appVersionId"`
	DownloadId   int    `json:"downloadId"`
	Url          string `json:"url"`
	Secret       string `json:"secret"`
}