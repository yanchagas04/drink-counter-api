package posts

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

type PostListResponse struct {
	Posts []PostResponse `json:"posts"`
}