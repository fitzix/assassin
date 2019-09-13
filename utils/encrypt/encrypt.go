package encrypt

import "github.com/matoous/go-nanoid"

func GetNanoId() string {
	id, _ := gonanoid.Nanoid(16)
	return id
}
