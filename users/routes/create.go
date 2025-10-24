package routes

import (
	"drink-counter-api/users/models"
	"drink-counter-api/users/schemas"
	"encoding/json"
	"log"
	"net/http"

	"gorm.io/gorm"
)

func CreateHandler(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var userRequest schemas.UserRequest
	err := json.NewDecoder(r.Body).Decode(&userRequest)
	log.Default().Println("userRequest -> ", userRequest)
	if err != nil {
		http.Error(w, "Invalid request body: " + err.Error(), http.StatusBadRequest)
		return
	}
	user := models.User {
		Name: userRequest.Name,
		Username: userRequest.Username,
		Email: userRequest.Email,
		Password: userRequest.Password,
	}
	log.Default().Println("user struct -> ", user)
	id := db.Create(&user)
	if id.Error != nil {
		switch id.Error {
		case gorm.ErrInvalidData:
			http.Error(w, "Invalid data: missing fields or invalid format", http.StatusBadRequest)
		case gorm.ErrDuplicatedKey:
			http.Error(w, "User already exists", http.StatusConflict)
		default:
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	response := schemas.UserResponse{
		Message: "User created successfully",
		Data: schemas.UserData{
			ID: user.ID,
			Name: user.Name,
			Username: user.Username,
			Email: user.Email,
			CreatedAt: user.CreatedAt.String(),
			UpdatedAt: user.UpdatedAt.String(),
		},
	}
	json.NewEncoder(w).Encode(response)
}