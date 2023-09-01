package util

import (
	gonnaId "github.com/matoous/go-nanoid/v2"
	"thichlab-backend-docs/constant"
)

// NewNanoIdString generate random string
func NewNanoIdString() string {
	id, err := gonnaId.Generate(constant.NanoIdAlphabet, constant.NanoIdSize)

	if err != nil {
		return ""
	}

	return id
}
