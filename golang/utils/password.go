package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func VerifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func HMAC(message string) string {
	key := []byte(os.Getenv("ENCRYPTION_AES_SECRET_KEY"))

	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	signature := h.Sum(nil)

	return hex.EncodeToString(signature)
}
