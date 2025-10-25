package utils

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"gorm.io/gorm"
)

func EntityNotFound(entity string) ErrorResponse {
	log.Default().Println("DB ERROR: " + entity + " not found")
	return ErrorResponse{
		Message: entity + " not found",
	}
}

func EntityAlreadyExists(entity string) ErrorResponse {
	log.Default().Println("DB ERROR: " + entity + " already exists")
	return ErrorResponse{
		Message: entity + " already exists",
	}
}

func SomethingWentWrongDB() ErrorResponse {
	log.Default().Println("DB ERROR: Something went wrong")
	return ErrorResponse{
		Message: "Something went wrong",
	}
}

func CheckDatabaseErrors(err error, w http.ResponseWriter, entity string) bool {
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(EntityNotFound(entity))
			return true
		} else if errors.Is(err, gorm.ErrDuplicatedKey) {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(EntityAlreadyExists(entity))
			return true
		} else if errors.Is(err, gorm.ErrInvalidData) {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(InvalidRequestBody(nil))
			return true
		} else if errors.Is(err, gorm.ErrInvalidField) {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(InvalidRequestBody(nil))
			return true
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(SomethingWentWrongDB())
		return true
	}
	return false

}