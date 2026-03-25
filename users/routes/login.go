package routes

import (
	"drink-counter-api/users/models"
	"drink-counter-api/users/schemas"
	UserUtils "drink-counter-api/users/utils"
	"drink-counter-api/utils"
	DatabaseErrors "drink-counter-api/utils/db_errors"
	SchemaErros "drink-counter-api/utils/schema_errors"
	"encoding/json"
	"log"
	"net/http"

	"gorm.io/gorm"
)

// Login using email and password, returning a access token.
func UserLoginHandler(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var userRequest schemas.UserLoginRequest
	err := json.NewDecoder(r.Body).Decode(&userRequest)
	log.Default().Println("userRequest -> ", userRequest)
	if SchemaErros.CheckSchemaErrors(err, w, userRequest) {
		return
	}
	var user models.User
	result := db.Where("email = ?", userRequest.Email).First(&user)
	if DatabaseErrors.CheckDatabaseErrors(result.Error, w, "User") {
		return
	}
	if UserUtils.WrongPassword(userRequest.Password, user.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(SchemaErros.ErrorResponse{
			Message: "Invalid credentials",
		})
		return
	}
	w.WriteHeader(http.StatusOK)
	response := schemas.UserLoginResponse {
		Message: "User logged in successfully",
		Token: UserUtils.GenerateToken(user),
		Data: schemas.UserData{
			ID: user.ID,
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