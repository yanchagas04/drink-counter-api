package routes

import (
	"drink-counter-api/users/models"
	"drink-counter-api/users/schemas"
	SchemaErrors "drink-counter-api/utils"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func GetHandler(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(SchemaErrors.ErrorResponse{
			Message: "ID is missing or is not an acceptable type",
		})
	}
	var user models.User
	result := db.First(&user, uint(id))
	if SchemaErrors.CheckDatabaseErrors(result.Error, w, "User") {
		return
	}
	w.WriteHeader(http.StatusOK)
	response := schemas.UserResponse{
		Message: "User found successfully",
		Data: schemas.UserData{
			ID:        user.ID,
			Name:      user.Name,
			Username:  user.Username,
			Email:     user.Email,
			CreatedAt: user.CreatedAt.String(),
			UpdatedAt: user.UpdatedAt.String(),
			DeletedAt: user.DeletedAt.Time.String(),
		},
	}
	json.NewEncoder(w).Encode(response)
}
