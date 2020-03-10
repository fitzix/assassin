package models

import "mime/multipart"

type ImageUploadReq struct {
	File *multipart.FileHeader `json:"file" form:"file" validate:"required"`
}

type Carousel struct {
	ID    uint   `json:"id"`
	AppId string `json:"appId"`
	Url   string `json:"url"`
}
