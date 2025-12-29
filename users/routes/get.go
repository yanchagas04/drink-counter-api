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

func GetHandler(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	q := r.URL.Query().Get("q")
	log.Default().Println("q -> ", q)
	var users []models.User
	var usersList []schemas.UserData
	result := db.Where("name LIKE ? OR username LIKE ?", q + "%", q + "%").Find(&users)
	if SchemaErrors.CheckDatabaseErrors(result.Error, w, "User") {
		return
	}
	w.WriteHeader(http.StatusOK)
	for _, user := range users {
		usersList = append(usersList, schemas.UserData {
			ID: user.ID,
			Name: user.Name,
			Username: user.Username,
			Email: user.Email,
			CreatedAt: user.CreatedAt.String(),
			UpdatedAt: user.UpdatedAt.String(),
			DeletedAt: user.DeletedAt.Time.String(),
		})
	}
	log.Default().Println("usersList -> ", usersList)
	response := schemas.UserListResponse{
		Message: "Users found successfully",
		Data: usersList,
	}
	json.NewEncoder(w).Encode(response)
}
