package routes

import (
	"drink-counter-api/users/models"
	DatabaseErrors "drink-counter-api/utils/db_errors"
	SchemaErrors "drink-counter-api/utils/schema_errors"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// Deletes a user (token required). Only the user who created the user can delete it's own user.
func DeleteUserHandler(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(SchemaErrors.ErrorResponse{
			Message: "Invalid id",
		})
		return
	}
	var user models.User
	result := db.Delete(&user, "id = ?", id)
	if DatabaseErrors.CheckDatabaseErrors(result.Error, w, "User") {
		return
	}
	if result.RowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(SchemaErrors.ErrorResponse{
			Message: "User not found",
		})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(SchemaErrors.ErrorResponse{
		Message: "User deleted successfully",
	})
}