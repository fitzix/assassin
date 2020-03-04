package models

import "github.com/dgrijalva/jwt-go"

type Token struct {
	Uid int `json:"uid"`
	// TODO 用于踢出用户
	Code uint `json:"code"`
}

type JwtClaims struct {
	Token
	jwt.StandardClaims
}
