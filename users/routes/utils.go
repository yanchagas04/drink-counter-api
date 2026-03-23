package routes

import (
	"drink-counter-api/users/models"
	DatabaseErrors "drink-counter-api/utils/db_errors"
	"log"

	"gorm.io/gorm"
)

func userNotFound(db *gorm.DB, id uint) {
	user := models.User{}
	result := db.Where("id = ?", id).First(&user)
	if DatabaseErrors.CheckDatabaseErrors(result.Error, nil, "User") {
		return
	}
	if result.RowsAffected == 0 {
		log.Default().Println("User not found")
	}

}