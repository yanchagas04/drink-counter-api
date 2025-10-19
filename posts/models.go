package posts

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title string
	Description string
}

type PostResponse struct {
	ID uint `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type PostRequest struct {
	Title string `json:"title"`
	Description string `json:"description"`
}