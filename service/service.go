package service

import (
	"encoding/base64"
	"fmt"

	"github.com/fitzix/assassin/utils/encrypt"
	"golang.org/x/crypto/scrypt"
)

func AesEncrypt(rawData string) string {
	// msg, err := encrypt.AESCBCPKCS7EncryptWithIV([]byte(appConf.Encrypt.Iv), []byte(appConf.Encrypt.Key), []byte(rawData))
	msg, err := encrypt.AESCBCPKCS7Encrypt([]byte(appConf.Encrypt.Key), []byte(rawData))
	if err != nil {
		zapLogger.Sugar().Errorf("AesEncrypt err: %s", err)
		return ""
	}

	return base64.StdEncoding.EncodeToString(msg)
}

func PassEncrypt(password string) string {
	dk, err := scrypt.Key([]byte(password), []byte(appConf.Salt), 32768, 8, 1, 16)
	if err != nil {
		zapLogger.Sugar().Errorf("PassEncrypt err: %s", err)
		return ""
	}
	return fmt.Sprintf("%x", dk)
}
