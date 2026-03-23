package utils

import (
	"crypto/sha256"
	"drink-counter-api/users/models"
	"encoding/hex"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Function to hash password
func HashPassword(password string) string {
	hash := sha256.Sum256([]byte(password))
	return hex.EncodeToString(hash[:])
}

// Function to verify password, returns true if password is correct else false
func VerifyPassword(password string, hashedPassword string) bool {
	return HashPassword(password) == hashedPassword
}

// Function to verify if password is wrong (just for abstraction), returns true if password is wrong else false
func WrongPassword(password string, hashedPassword string) bool {
	return !VerifyPassword(password, hashedPassword)
}

// Function to generate a token
func GenerateToken(user models.User) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"email":    user.Email,
		"id":       user.ID,
		"name":     user.Name,
		"password": user.Password,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, _ := token.SignedString([]byte("secret"))
	return tokenString
}

// Function to validate a token
func ValidateToken(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte("secret"), nil
	})
	if err != nil {
		return false, err
	}
	return token.Valid, nil
}

// Function to get id from token
func GetIdFromToken(tokenString string) (uint, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte("secret"), nil
	})
	if err != nil {
		return 0, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		id, ok := claims["id"].(float64)
		if !ok {
			return 0, jwt.ErrTokenInvalidClaims
		}
		return uint(id), nil
	}
	return 0, jwt.ErrTokenInvalidId
}