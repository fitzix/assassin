package models

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`
}

type AppCategory struct {
	AppID      string `json:"-"`
	CategoryID int    `json:"id"`
}
