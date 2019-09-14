package model

type UserLogin struct {
	UserName string `json:"userName" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
