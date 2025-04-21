package utils

import (
	gonanoid "github.com/matoous/go-nanoid/v2"
)

func GenerateID() string {
	id, err := gonanoid.Generate("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz", 7)
	if err != nil {
		return ""
	}
	return id
}
