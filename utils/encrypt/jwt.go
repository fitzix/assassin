package encrypt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/fitzix/assassin/service"
	"github.com/fitzix/assassin/service/model"
	"github.com/fitzix/assassin/utils"
)



type asnJwtClaims struct {
	model.Token
	jwt.StandardClaims
}

func GenJwt(t model.Token) (string, error) {
	conf := utils.GetConf()
	claims := &asnJwtClaims{
		Token:          t,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(conf.Jwt.Expires)).Unix(),
			Issuer:    conf.Jwt.Issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(conf.Jwt.Secret))
}

func ParseToken(tokenString string) (*model.Token, error) {
	secret := utils.GetConf().Jwt.Secret
	token, err := jwt.ParseWithClaims(tokenString, &asnJwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, service.ErrInvalidToken
	}

	if claims, ok := token.Claims.(*asnJwtClaims); ok && token.Valid {
		return &claims.Token, nil
	}
	return nil, service.ErrInvalidToken
}
