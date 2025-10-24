package users

type UserData struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UserRequest struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponse struct {
	Message string   `json:"message"`
	Data    UserData `json:"data"`
}

type UserListResponse struct {
	Message string     `json:"message"`
	Data    []UserData `json:"data"`
}