package posts

type CommentData struct {
	ID        uint   `json:"id"`
	Content   string `json:"content"`
	Author    uint   `json:"author"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type PostData struct {
	ID          uint          `json:"id"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Likes       int           `json:"likes"`
	CreatedAt   string        `json:"created_at"`
	UpdatedAt   string        `json:"updated_at"`
}

type PostRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type PostResponse struct {
	Message string   `json:"message"`
	Data    PostData `json:"data"`
}

type PostListResponse struct {
	Message string     `json:"message"`
	Data    []PostData `json:"data"`
}