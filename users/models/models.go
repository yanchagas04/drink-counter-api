package models

import (
	posts "drink-counter-api/posts/models"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID uint
	Name string
	Username string `gorm:"unique"`
	Email string `gorm:"unique"`
	Password string
	Posts []posts.Post `gorm:"foreignKey:Author;OnDelete:CASCADE;OnUpdate:CASCADE"`
	Comments []posts.Comment `gorm:"foreignKey:Author;OnDelete:CASCADE;OnUpdate:CASCADE"`
}