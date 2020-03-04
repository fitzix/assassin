package utils

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/fitzix/assassin/models"
)

var (
	ErrInvalidToken = errors.New("jwt: invalid jwt key")
)

func GenJwt(t models.Token, conf models.Jwt) (string, error) {
	claims := &models.JwtClaims{
		Token: t,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(conf.Expires)).Unix(),
			Issuer:    conf.Issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(conf.Secret))
}

func ParseToken(tokenString string, conf models.Jwt) (*models.Token, error) {
	secret := conf.Secret
	token, err := jwt.ParseWithClaims(tokenString, &models.JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, ErrInvalidToken
	}

	if claims, ok := token.Claims.(*models.JwtClaims); ok && token.Valid {
		return &claims.Token, nil
	}
	return nil, ErrInvalidToken
}
