package models

type Hot struct {
	AppId uint64 `json:"-"`
	Hot   int    `json:"hot"`
	View  int    `json:"view"`
}
