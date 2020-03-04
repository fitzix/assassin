package utils

import "github.com/matoous/go-nanoid"

func GenNanoId() string {
	id, _ := gonanoid.Nanoid(16)
	return id
}
