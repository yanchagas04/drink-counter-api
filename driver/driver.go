package driver

import (
	"drink-counter-api/utils"
	"errors"
	"net/http"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	utils.LoadEnv()
	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{TranslateError: true})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}

func Close(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to close database")
	}
	sqlDB.Close()
}

type ErrorResponse struct {
	Message string `json:"message"`
}

type DriverError struct {
	Code    int
	Response ErrorResponse
}

func CheckErrors(result *gorm.DB, entity string) DriverError {
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return DriverError{
				Code: http.StatusNotFound,
				Response: ErrorResponse{
					Message: entity + " not found",
				},
			}
		}
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return DriverError{
				Code: http.StatusBadRequest,
				Response: ErrorResponse{
					Message: entity + " already exists",
				},
			}
		}
		if errors.Is(result.Error, gorm.ErrInvalidField) {
			return DriverError{
				Code: http.StatusBadRequest,
				Response: ErrorResponse{
					Message: entity + " invalid field: " + result.Error.Error(),
				},
			}
		}
		if errors.Is(result.Error, gorm.ErrForeignKeyViolated) {
			return DriverError{
				Code: http.StatusBadRequest,
				Response: ErrorResponse{
					Message: entity + " invalid field: " + result.Error.Error(),
				},
			}
		}
		if errors.Is(result.Error, gorm.ErrInvalidData) {
			return DriverError{
				Code: http.StatusBadRequest,
				Response: ErrorResponse{
					Message: entity + " invalid data: " + result.Error.Error(),
				},
			}
		}
		return DriverError{
			Code: http.StatusInternalServerError,
			Response: ErrorResponse{
				Message: entity + " error: " + result.Error.Error(),
			},
		}
	}
	return DriverError{}
}