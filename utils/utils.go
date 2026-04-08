package utils

import (
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var DATEFORMAT = "2006-01-02T15:04:05Z07:00"

var PAGESIZE = 10

func CalculateOffset(page int) int {
	return (page - 1) * PAGESIZE
}

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}
}

func VerifyIfDeleted(field gorm.DeletedAt) string {
	if field.Valid {
		return field.Time.Format(DATEFORMAT)
	}
	return ""
}