package routes

import (
	"drink-counter-api/users/models"
	"drink-counter-api/users/schemas"
	UserUtils "drink-counter-api/users/utils"
	"drink-counter-api/utils"
	DatabaseErrors "drink-counter-api/utils/db_errors"
	SchemaErrors "drink-counter-api/utils/schema_errors"
	"encoding/json"
	"log"
	"net/http"

	"gorm.io/gorm"
)

// Create a new user.
func CreateUserHandler(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
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
		Password: UserUtils.HashPassword(userRequest.Password),
	}
	log.Default().Println("user struct -> ", user)
	result := db.Create(&user)
	if DatabaseErrors.CheckDatabaseErrors(result.Error, w, "User"){
		return
	}
	w.WriteHeader(http.StatusCreated)
	response := schemas.UserResponse{
		Message: "User created successfully",
		Data: schemas.UserData{
			Name: user.Name,
			Username: user.Username,
			Email: user.Email,
			CreatedAt: user.CreatedAt.Format(utils.DATEFORMAT),
			UpdatedAt: user.UpdatedAt.Format(utils.DATEFORMAT),
			DeletedAt: utils.VerifyIfDeleted(user.DeletedAt),
		},
	}
	json.NewEncoder(w).Encode(response)
}
