package service

import (
	"encoding/json"
	"fmt"
	"mime/multipart"
	"path"
	"path/filepath"

	"github.com/fitzix/assassin/consts"
	"github.com/fitzix/assassin/models"
	"github.com/fitzix/assassin/utils"
	"github.com/minio/minio-go/v6"
)

// 检查图片资源是否公开
func checkAndSetBucketPolicy() {
	p, err := s3.GetBucketPolicy(conf.Bucket)
	if err != nil {
		logger.Fatalf("s3 get bucket policy err: %s", err)
		return
	}
	if p == "" {
		setS3Policy(consts.S3PolicyAllowImageStatic)
		return
	}

	var policy models.S3Policy
	if err := json.Unmarshal([]byte(p), &policy); err != nil {
		logger.Fatalf("s3 parse bucket policy err: %s", err)
		return
	}
	if len(policy.Statement) > 0 {
		for _, v := range policy.Statement {
			if v.Sid == "AllowImageStatic" {
				logger.Info("s3 bucket policy checked ok")
				return
			}
		}
	}

	var initPolicy models.S3Policy
	if err := json.Unmarshal([]byte(consts.S3PolicyAllowImageStatic), &initPolicy); err != nil {
		logger.Fatalf("s3 parse init bucket policy err: %s", err)
		return
	}

	policy.Statement = append(policy.Statement, initPolicy.Statement[0])

	b, err := json.Marshal(&policy)
	if err != nil {
		logger.Fatalf("s3 marshal new bucket policy err: %s", err)
		return
	}
	setS3Policy(string(b))
}

func setS3Policy(policy string) {
	if err := s3.SetBucketPolicy(conf.Bucket, fmt.Sprintf(policy, conf.Bucket)); err != nil {
		logger.Fatalf("set s3 bucket policy err: %s", err)
	}
	logger.Infof("successfully set bucket policy %s", conf.Bucket)
}

func PutImage(file *multipart.FileHeader) (string, error) {
	fileName := utils.GenNanoId() + filepath.Ext(file.Filename)
	filePath := path.Join(conf.ImgPrefix, fileName)
	f, err := file.Open()
	if err != nil {
		return "", err
	}
	defer f.Close()
	if _, err := s3.PutObject(conf.Bucket, filePath, f, file.Size, minio.PutObjectOptions{}); err != nil {
		return "", err
	}
	return fileName, nil
}
