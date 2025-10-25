package routes

import (
	"drink-counter-api/users/models"
	"drink-counter-api/users/schemas"
	SchemaErrors "drink-counter-api/utils"
	"encoding/json"
	"log"
	"net/http"

	"gorm.io/gorm"
)

func CreateHandler(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var userRequest schemas.UserRequest
	err := json.NewDecoder(r.Body).Decode(&userRequest)
	log.Default().Println("userRequest -> ", userRequest)
	if SchemaErrors.CheckSchemaErrors(err, w, userRequest){
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
	if SchemaErrors.CheckDatabaseErrors(result.Error, w, "User"){
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
