package service

import (
	"encoding/base64"

	"github.com/fitzix/assassin/utils"
)

func AesEncrypt(rawData string) string {
	// msg, err := encrypt.AESCBCPKCS7EncryptWithIV([]byte(appConf.Encrypt.Iv), []byte(appConf.Encrypt.Key), []byte(rawData))
	msg, err := utils.AESCBCPKCS7Encrypt([]byte(conf.Encrypt.Key), []byte(rawData))
	if err != nil {
		logger.Errorf("AesEncrypt err: %s", err)
		return ""
	}

	return base64.StdEncoding.EncodeToString(msg)
}
