package routes

import (
	"drink-counter-api/users/models"
	"drink-counter-api/users/schemas"
	"drink-counter-api/utils"
	Utils "drink-counter-api/utils"
	DatabaseErrors "drink-counter-api/utils/db_errors"
	SchemaErrors "drink-counter-api/utils/schema_errors"
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
		Password: Utils.HashPassword(userRequest.Password),
	}
	log.Default().Println("user struct -> ", user)
	result := db.Save(&user)
	if DatabaseErrors.CheckDatabaseErrors(result.Error, w, "User") {
		return
	}
	w.WriteHeader(http.StatusOK)
	response := schemas.UserResponse {
		Message: "User updated successfully",
		Data: schemas.UserData {
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