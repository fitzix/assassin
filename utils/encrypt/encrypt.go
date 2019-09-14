package encrypt

import (
	"fmt"

	"github.com/fitzix/assassin/utils"
	"github.com/matoous/go-nanoid"
	"golang.org/x/crypto/scrypt"
)

func GetNanoId() string {
	id, _ := gonanoid.Nanoid(16)
	return id
}

func PassEncrypt(password string) string {
	salt := utils.GetConf().Salt
	dk, err := scrypt.Key([]byte(password), []byte(salt), 32768, 8, 1, 16)
	if err != nil {
		utils.GetLogger().Sugar().Errorf("password encrypt err: %s", err)
	}
	return fmt.Sprintf("%x", dk)
}
