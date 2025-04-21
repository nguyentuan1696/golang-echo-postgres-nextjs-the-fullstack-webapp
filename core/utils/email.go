package utils

import (
	"net"
	"regexp"
	"strings"
)

var (
	// RFC 5322 compliant email regex
	emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
)

// IsValidEmail checks if email format is valid and within length limits
func IsValidEmail(email string) bool {
	if len(email) > 254 {
		return false
	}
	return emailRegex.MatchString(email)
}

// IsValidEmailDomain checks if email domain has valid MX records
func IsValidEmailDomain(email string) bool {
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return false
	}

	domain := parts[1]
	if len(domain) > 253 {
		return false
	}

	mx, err := net.LookupMX(domain)
	if err != nil || len(mx) == 0 {
		return false
	}

	return true
}