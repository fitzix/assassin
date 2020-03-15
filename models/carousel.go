package models

import "mime/multipart"

type ImageUploadReq struct {
	File *multipart.FileHeader `json:"file" form:"file" validate:"required"`
}

type Carousel struct {
	ID    uint   `json:"-"`
	AppID uint64 `json:"-"`
	Url   string `json:"url"`
}
