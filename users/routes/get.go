package routes

import (
	"drink-counter-api/users/models"
	"drink-counter-api/users/schemas"
	"drink-counter-api/utils"
	DatabaseErrors "drink-counter-api/utils/db_errors"
	SchemaErrors "drink-counter-api/utils/schema_errors"
	"encoding/json"
	"log"
	"net/http"

	"gorm.io/gorm"
)

func GetHandler(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	q := r.URL.Query().Get("q")
	log.Default().Println("q -> ", q)
	var users []models.User
	var usersList []schemas.UserData
	
	result := db.Where("name LIKE ? OR username LIKE ?", q + "%", q + "%").Find(&users)
	if DatabaseErrors.CheckDatabaseErrors(result.Error, w, "User") {
		return
	}
	if len(users) == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(SchemaErrors.ErrorResponse{
			Message: "No users found",
		})
		return
	}
	w.WriteHeader(http.StatusOK)
	for _, user := range users {
		usersList = append(usersList, schemas.UserData {
			Name: user.Name,
			Username: user.Username,
			Email: user.Email,
			CreatedAt: user.CreatedAt.Format(utils.DATEFORMAT),
			UpdatedAt: user.UpdatedAt.Format(utils.DATEFORMAT),
			DeletedAt: utils.VerifyIfDeleted(user.DeletedAt),
		})
	}
	log.Default().Println("usersList -> ", usersList)
	response := schemas.UserListResponse{
		Message: "Users found successfully",
		Data: usersList,
	}
	json.NewEncoder(w).Encode(response)
}
