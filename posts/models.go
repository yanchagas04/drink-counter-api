package posts

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	ID uint
	Content string
	PostID uint
	Author uint
}

type Post struct {
	gorm.Model
	ID uint
	Title string
	Description string
	Amount int
	Likes int
	Comments []Comment `gorm:"foreignKey:PostID;OnDelete:CASCADE;OnUpdate:CASCADE"`
	Author uint
}
