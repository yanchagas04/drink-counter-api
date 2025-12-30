package utils

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}
}

func HashPassword(password string) string {
	hash := sha256.Sum256([]byte(password))
	return hex.EncodeToString(hash[:])
}

func VerifyPassword(password string, hashedPassword string) bool {
	return HashPassword(password) == hashedPassword
}