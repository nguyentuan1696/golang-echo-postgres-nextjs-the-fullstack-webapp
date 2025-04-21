package utils

import (
	"strings"

	"github.com/gosimple/slug"
)

func GenerateIDAndSlug(s string) (string, string) {
	id := GenerateID()
	if s == "" {
		return id, id
	}

	var b strings.Builder
	b.WriteString(slug.Make(s))
	b.WriteByte('-')
	b.WriteString(id)
	return id, b.String()
}
