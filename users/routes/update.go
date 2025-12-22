package routes

import (
	"drink-counter-api/users/models"
	"drink-counter-api/users/schemas"
	SchemaErrors "drink-counter-api/utils"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func UpdateHandler(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var userRequest schemas.UserRequest
	err := json.NewDecoder(r.Body).Decode(&userRequest)
	log.Default().Println("userRequest -> ", userRequest)
	if SchemaErrors.CheckSchemaErrors(err, w, userRequest){
		return
	}
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	user := models.User {
		ID: uint(id),
		Name: userRequest.Name,
		Username: userRequest.Username,
		Email: userRequest.Email,
		Password: userRequest.Password,
	}
	log.Default().Println("user struct -> ", user)
	result := db.Save(&user)
	if SchemaErrors.CheckDatabaseErrors(result.Error, w, "User") {
		return
	}
	w.WriteHeader(http.StatusOK)
	response := schemas.UserResponse {
		Message: "User updated successfully",
		Data: schemas.UserData {
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