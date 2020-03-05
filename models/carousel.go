package models

import "mime/multipart"

type ImageUploadReq struct {
	File *multipart.FileHeader `json:"file" form:"file" validate:"required"`
}
