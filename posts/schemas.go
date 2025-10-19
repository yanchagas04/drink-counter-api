package posts

type PostRequest struct {
	Title string `json:"title"`
	Description string `json:"description"`
}

type PostData struct {
	ID uint `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type PostResponse struct {
	Message string `json:"message"`
	Data PostData `json:"data"`
}

type PostListResponse struct {
	Message string `json:"message"`
	Data []PostData `json:"data"`
}