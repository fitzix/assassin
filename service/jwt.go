package service

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type asnJwtClaims struct {
	Token
	jwt.StandardClaims
}

func GenJwt(t Token) (string, error) {
	conf := appConf.Jwt
	claims := &asnJwtClaims{
		Token: t,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(conf.Expires)).Unix(),
			Issuer:    conf.Issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(conf.Secret))
}

func ParseToken(tokenString string) (*Token, error) {
	secret := appConf.Jwt.Secret
	token, err := jwt.ParseWithClaims(tokenString, &asnJwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, ErrInvalidToken
	}

	if claims, ok := token.Claims.(*asnJwtClaims); ok && token.Valid {
		return &claims.Token, nil
	}
	return nil, ErrInvalidToken
}
