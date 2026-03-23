package utils

import (
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var DATEFORMAT = "2006-01-02T15:04:05Z07:00"

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