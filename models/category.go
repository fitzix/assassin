package models

type CategoryCreated struct {
	Name string `json:"name"`
	Icon string `json:"icon"`
}

type Category struct {
	ID int `json:"id"`
	CategoryCreated
}

type AppCategory struct {
	AppID      uint64 `json:"-"`
	CategoryID int    `json:"id"`
}
