package models

import "github.com/fitzix/assassin/ent"

type User struct {
	ID       int    `json:"-"`
	UID      string `json:"uid"`
	Name     string `json:"name" binding:"required,email"`
	Password string `json:"-"`
	RoleId   int    `json:"roleId"`
	Code     uint   `json:"-"`
	Status   bool   `json:"status"`
}

type UserLoginReq struct {
	UserName string `json:"user" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UserCreateReq struct {
	UserLoginReq
	RoleID int `json:"role" validate:"required"`
}

type UserLoginRsp struct {
	*ent.User
	Token string `json:"token"`
}
