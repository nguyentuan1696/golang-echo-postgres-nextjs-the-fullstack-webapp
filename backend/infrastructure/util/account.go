package util

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"regexp"
	"strings"
)

var emailPattern = regexp.MustCompile(`[a-zA-Z.\-_][a-zA-Z.\-_0-9]{4,}@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z0-9]+\.)+[a-zA-Z]{2,}))`)

func isEmail(s string) bool {
	// SECURITY: the len() check prevents a regex ddos via overly large usernames
	return len(s) < 255 && emailPattern.MatchString(s)
}

func EmailValidator(email string) error {
	if email == "" {
		return fmt.Errorf("Email Empty")
	}

	if !isEmail(email) {
		return fmt.Errorf("Invalid Email")
	}
	return nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func PasswordValidator(password string) error {
	if password == "" {
		return fmt.Errorf("Password Empty")
	}

	if len(password) < 6 {
		return fmt.Errorf("Password too short")
	}

	re := regexp.MustCompile(`^*[0-9]`)
	if !re.MatchString(password) {
		return fmt.Errorf("Password Must Contain Number")
	}

	re = regexp.MustCompile(`^*[a-zA-Z]`)
	if !re.MatchString(password) {
		return fmt.Errorf("Password Must Contain Letter")
	}

	re = regexp.MustCompile(`^[a-zA-Z0-9 !"#$%&'()*+,-./:;<=>?@[\]^_{|}]*$`)
	if !re.MatchString(password) {
		return fmt.Errorf("Invalid Character found")
	}

	return nil
}

func UsernameValidator(username string) error {
	if len(username) < 4 {
		return fmt.Errorf("Username too short")
	}

	if len(username) > 32 {
		return fmt.Errorf("Username too long")
	}

	re := regexp.MustCompile(`^[a-zA-Z0-9]+(?:[_.-][a-zA-Z0-9]+)*$`)
	if !re.MatchString(username) {
		return fmt.Errorf("Username Wrong Format")
	}
	if strings.Contains(username, " ") {
		return fmt.Errorf("Username cannot have spaces")
	}
	return nil
}
