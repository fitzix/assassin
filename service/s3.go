package service

import (
	"mime/multipart"
	"path/filepath"

	"github.com/fitzix/assassin/utils"
	"github.com/minio/minio-go/v6"
)

func PutImage(file *multipart.FileHeader) (string, error) {
	fileName := "images/" + utils.GenNanoId() + filepath.Ext(file.Filename)
	f, err := file.Open()
	if err != nil {
		return "", err
	}
	defer f.Close()
	if _, err := s3.PutObject(conf.Bucket, fileName, f, file.Size, minio.PutObjectOptions{}); err != nil {
		return "", err
	}
	return conf.Endpoint + fileName, nil
}
