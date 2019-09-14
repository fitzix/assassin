package models

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name" binding:"required,email"`
	Password string `json:"-"`
	RoleId   int    `json:"roleId"`
	Code     uint   `json:"-"`
}
