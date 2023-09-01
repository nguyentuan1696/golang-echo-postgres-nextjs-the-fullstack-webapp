package util

import (
	"github.com/gosimple/slug"
	"strings"
)

func ToSlugWithId(slug, id string) string {
	var sb strings.Builder
	sb.WriteString(slug)
	sb.WriteString("-")
	sb.WriteString(id)
	return sb.String()
}

func ToSlugFromTitleWithoutId(title string) string {
	var sb strings.Builder
	sb.WriteString(title)
	sb.WriteString("-")
	return sb.String()
}

func ToSlugFromTitleWithId(title, id string) string {
	var sb strings.Builder
	text := slug.Make(title)
	sb.WriteString(text)
	sb.WriteString("-")
	sb.WriteString(id)
	return sb.String()
}

func ToSlug(title string) string {
	return slug.Make(strings.TrimSpace(title))
}
