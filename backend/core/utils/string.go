package utils

import (
	"strings"

	"github.com/google/uuid"
)

// TrimSpace removes leading, trailing, and multiple spaces between words
func TrimSpace(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

// TrimSpacePointer handles string pointer and returns trimmed string pointer
func TrimSpacePointer(s *string) *string {
	if s == nil {
		return nil
	}
	trimmed := TrimSpace(*s)
	return &trimmed
}

// TrimAllSpaces removes all spaces from a string
func TrimAllSpaces(s string) string {
	return strings.ReplaceAll(s, " ", "")
}

// IsEmpty checks if a string is empty after trimming spaces
func IsEmpty(s string) bool {
	return TrimSpace(s) == ""
}

// ToNumber converts string to number, returns 0 if conversion fails
func ToNumber(s string) int {
	// Remove all spaces
	s = TrimAllSpaces(s)

	// Convert string to number
	var result int
	for _, ch := range s {
		// Check if character is digit
		if ch < '0' || ch > '9' {
			return 0
		}
		// Build number
		result = result*10 + int(ch-'0')
	}

	return result
}

// ToNumberWithDefault converts string to number with a default value
func ToNumberWithDefault(s string, defaultValue int) int {
	if IsEmpty(s) {
		return defaultValue
	}
	return ToNumber(s)
}

func ToString(s uuid.UUID) string {
	return s.String()
}
