package models

import "github.com/dgrijalva/jwt-go/v4"

type Token struct {
	Uid string `json:"uid"`
	// TODO 用于踢出用户
	Code uint `json:"code"`
}

type JwtClaims struct {
	Token
	jwt.StandardClaims
}
