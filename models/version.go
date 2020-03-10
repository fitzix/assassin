package models

import (
	"time"
)

type Version struct {
	ID        int       `json:"id"`
	Name      string    `json:"name" binding:"required"`
	AppId     string    `json:"appId"`
	Size      string    `json:"size"`
	Status    bool      `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
}


