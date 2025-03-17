package utils

import (
	"math/rand"
	"strings"
	"time"
	"unicode"
)

// Convert a camelCase string to PascalCase
func CamelToPascalCase(s string) string {
	if !IsCamelCase(s) {
		return s
	}
	return strings.ToUpper(string(s[0])) + s[1:]
}

func ToUpper(s string) string {
	if s == "" {
		return s
	}

	return strings.ToUpper(string(s[0])) + s[1:]
}

func PointerString(s string) *string {
	if s == "" {
		return nil
	}

	return &s
}

func IsCamelCase(s string) bool {
	if s == "" {
		return false
	}
	// Check if the first character is lowercase
	if !unicode.IsLower(rune(s[0])) {
		return false
	}
	// Check if there are any spaces or underscores
	for _, r := range s {
		if r == ' ' || r == '_' {
			return false
		}
	}
	return true
}

func PascalCaseToCamelCase(s string) string {
	if IsCamelCase(s) {
		return s
	}

	return strings.ToLower(string(s[0])) + s[1:]
}

func RandomAlphanumeric(length int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))
	b := make([]byte, length)
	for i := range b {
		b[i] = letterBytes[r.Intn(len(letterBytes))]
	}

	return string(b)
}
