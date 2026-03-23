package routes

import (
	"fmt"
	"net/http"

	"drink-counter-api/users/models"
	"drink-counter-api/users/schemas"
	UserServices "drink-counter-api/users/services"
	"drink-counter-api/utils"
	DatabaseErrors "drink-counter-api/utils/db_errors"
	"encoding/json"

	"gorm.io/gorm"
)

// Get a user's public profile and it's id (token required). Only the user who created the user can access.
func GetUserHandler(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // seta o header da resposta como json
	authorization := UserServices.AuthService(w, r.Header.Get("Authorization"))
	if !authorization.Valid {
		return
	}
	fmt.Println("id = ", *(authorization.UserId))
	var user models.User
	id := *(authorization.UserId)
	result := db.First(&user, "id = ?", id)
	if DatabaseErrors.CheckDatabaseErrors(result.Error, w, "User") { // verifica erros no banco
		return
	}
	response := schemas.UserData{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
		Username: user.Username,
		CreatedAt: user.CreatedAt.Format(utils.DATEFORMAT),
		UpdatedAt: user.UpdatedAt.Format(utils.DATEFORMAT),
		DeletedAt: utils.VerifyIfDeleted(user.DeletedAt),
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(schemas.UserResponse{
		Message: "User found successfully",
		Data: response,
	})
}