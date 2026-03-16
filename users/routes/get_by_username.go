package routes

import (
	"fmt"
	"net/http"

	"drink-counter-api/users/models"
	"drink-counter-api/users/schemas"
	"drink-counter-api/utils"
	DatabaseErrors "drink-counter-api/utils/db_errors"
	SchemaErrors "drink-counter-api/utils/schema_errors"
	"encoding/json"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func GetByUsername(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // seta o header da resposta como json
	username := mux.Vars(r)["username"] // pega o parâmetro username na rota
	if username == ""  { // verifica se o username está vazio
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(SchemaErrors.ErrorResponse{
			Message: "Username is missing or is not an acceptable type",
		})
		return
	}
	fmt.Println("username = " + username)
	var user models.User
	result := db.Where("username = ?", username).First(&user)
	if DatabaseErrors.CheckDatabaseErrors(result.Error, w, "User") { // verifica erros no banco
		return
	}
	response := schemas.UserData{
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