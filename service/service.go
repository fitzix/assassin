package service

import (
	"fmt"

	"github.com/fitzix/assassin/utils/encrypt"
	"golang.org/x/crypto/scrypt"
)

func AesEncrypt(rawData string) string {
	msg, err := encrypt.AesEncrypt(rawData, appConf.Encrypt.Key, appConf.Encrypt.Iv)
	if err != nil {
		zapLogger.Sugar().Errorf("AesEncrypt err: %s", err)
		return ""
	}
	return msg
}

func PassEncrypt(password string) string {
	dk, err := scrypt.Key([]byte(password), []byte(appConf.Salt), 32768, 8, 1, 16)
	if err != nil {
		zapLogger.Sugar().Errorf("PassEncrypt err: %s", err)
		return ""
	}
	return fmt.Sprintf("%x", dk)
}
