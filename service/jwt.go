package service

import (
	"github.com/fitzix/assassin/models"
	"github.com/fitzix/assassin/utils"
)

func GenJwt(t models.Token) (string, error) {
	return utils.GenJwt(t, conf.Jwt)
}

func ParseToken(t string) (*models.Token, error) {
	return utils.ParseToken(t, conf.Jwt)
}
