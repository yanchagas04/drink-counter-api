package utils

import (
	"crypto/sha256"
	"encoding/hex"

	"drink-counter-api/users/models"

	"github.com/golang-jwt/jwt/v5"
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

func GenerateToken(user models.User) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"email":    user.Email,
		"id":       user.ID,
		"name":     user.Name,
		"password": user.Password,
	})
	tokenString, _ := token.SignedString([]byte("secret"))
	return tokenString
}

func ValidateToken(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, nil
		}
		return []byte("secret"), nil
	})
	if err != nil {
		return false, err
	}
	if !token.Valid {
		return false, nil
	}
	return true, nil
}