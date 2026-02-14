package accounts

import (
	"strings"
	"unicode/utf8"
)

const (
	maxNameLength = 100
)

func NormalizeEmail(email string) string {
	return strings.TrimSpace(strings.ToLower(email))
}

func ValidateEmail(email string) error {
	if email == "" {
		return ErrInvalidInput
	}
	if !strings.Contains(email, "@") {
		return ErrInvalidInput
	}
	return nil
}

func ValidateName(name string) error {
	name = strings.TrimSpace(name)
	if name == "" {
		return ErrInvalidInput
	}
	if utf8.RuneCountInString(name) > maxNameLength {
		return ErrInvalidInput
	}
	return nil
}

func ValidatePassword(password string) error {
	if len(password) < 8 {
		return ErrInvalidInput
	}
	return nil
}
