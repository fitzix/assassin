package models

type Tag struct {
	ID   int    `json:"id"`
	Name string `json:"name" binding:"required"`
}

type AppTag struct {
	AppID uint64 `json:"-"`
	TagID uint   `json:"id"`
}
