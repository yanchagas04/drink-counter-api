package utils

import (
	"github.com/joho/godotenv"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}
}