package routes

import (
	"drink-counter-api/driver"
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
	result := db.Create(&user)
	w.Header().Set("Content-Type", "application/json")
	error := driver.CheckErrors(result, "User")
	if error.Code != 0 {
		log.Default().Println(error.Response.Message)
		w.WriteHeader(error.Code)
		json.NewEncoder(w).Encode(error.Response)
		return	
	}
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
			DeletedAt: user.DeletedAt.Time.String(),
		},
	}
	json.NewEncoder(w).Encode(response)
}
