package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword takes a plain text password and returns a bcrypt hashed version
// The cost parameter determines how computationally expensive the hash will be
// Default cost is 10, but you can adjust based on your security requirements
func HashPassword(password string) (string, error) {
	// Generate a bcrypt hash with default cost (10)
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedBytes), nil
}

// ComparePassword compares a hashed password with a plain text password
// Returns nil if they match, error otherwise
func ComparePassword(hashedPassword, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil
}
