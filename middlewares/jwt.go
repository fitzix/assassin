package middlewares

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type AsnClaims struct {
	Uid string `json:"uid"`
	jwt.StandardClaims
}

func Jwt(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// token := strings.Trim(c.GetHeader("Authorization"), "Bear ")

	}
}