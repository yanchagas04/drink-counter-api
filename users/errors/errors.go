package errors

import (
	"drink-counter-api/utils"
	"errors"

	"github.com/go-playground/validator/v10"
)

var Validator = validator.New()

func FormatValidationErrors(err error) map[string]string {
	errorsMap := make(map[string]string)

	var validationErrors validator.ValidationErrors
	if errors.As(err, &validationErrors) {
		for _, fieldError := range validationErrors {
			errorsMap[fieldError.Field()] = "Digite um " + fieldError.Tag() + " válido"
		}
	}
	return errorsMap
}

func UserNotFound() utils.ErrorResponse {
	return utils.ErrorResponse{
		Message: "User not found",
	}
}

func UserAlreadyExists() utils.ErrorResponse {
	return utils.ErrorResponse{
		Message: "User already exists",
	}
}

func InvalidField(msg string) utils.ErrorResponse {
	return utils.ErrorResponse{
		Message: msg,
	}
}