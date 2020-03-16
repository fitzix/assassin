package models

import (
	"time"
)

type Version struct {
	ID        int       `json:"id"`
	Name      string    `json:"name" binding:"required"`
	AppId     uint64    `json:"appId"`
	Size      uint64    `json:"size"`
	Status    bool      `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
	Sources   []Source  `json:"sources"`
}
