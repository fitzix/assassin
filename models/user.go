package models

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name" binding:"required,email"`
	Password string `json:"-"`
	RoleId   int    `json:"roleId"`
	Code     uint   `json:"-"`
	Status   bool   `json:"status"`
}

type UserLoginReq struct {
	UserName string `json:"name" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UserLoginRsp struct {
	Token string `json:"token"`
}
