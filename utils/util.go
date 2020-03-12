package utils

import (
	"time"

	"github.com/matoous/go-nanoid"
	"github.com/sony/sonyflake"
)

var flake = sonyflake.NewSonyflake(sonyflake.Settings{
	StartTime: time.Date(2020, 3, 3, 3, 3, 3, 3, time.UTC),
})

func GenNanoId() string {
	id, _ := gonanoid.Nanoid(16)
	return id
}

func NextID() uint64 {
	id, err := flake.NextID()
	if err != nil {
		panic(err)
	}
	return id
}
