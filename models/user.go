package models

type Login struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type LoginResp struct {
	Token string `json:"token"`
}
