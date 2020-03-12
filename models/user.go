package models

type User struct {
	ID       string `json:"-"`
	Name     string `json:"name"`
	Password string `json:"-"`
	RoleId   int    `json:"roleId"`
	Code     uint   `json:"-"`
	Status   bool   `json:"status"`
}

type UserLoginReq struct {
	UserName string `json:"user" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserCreateReq struct {
	UserLoginReq
	RoleID int `json:"role" binding:"required"`
}

type UserLoginRsp struct {
	User
	Token string `json:"token"`
}
