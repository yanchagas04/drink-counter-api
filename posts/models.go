package posts

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	ID uint
	Content string
}

type Post struct {
	gorm.Model
	ID uint
	Title string
	Description string
	Likes int
	Comments []Comment
}
