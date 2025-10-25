package utils

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

// Invalid type for field x
func InvalidField(field string) ErrorResponse {
	log.Default().Println("SCHEMA ERROR: Invalid type for field: " + field)
	return ErrorResponse{
		Message: "Invalid type for field: " + field,
	}
}

// Validation error for field x: msg
func FieldValidationError(msg string) ErrorResponse {
	log.Default().Println("SCHEMA ERROR: " + msg)
	return ErrorResponse{
		Message: msg,
	}
}

// Invalid request body
func InvalidRequestBody(msg *string) ErrorResponse {
	if msg == nil {
		log.Default().Println("SCHEMA ERROR: Invalid request body")
		return ErrorResponse{
			Message: "Invalid request body",
		}
	}
	log.Default().Println("SCHEMA ERROR: Invalid request body: " + *msg)
	return ErrorResponse{
		Message: "Invalid request body: " + *msg,
	}
}

// Something went wrong
func SomethingWentWrongSchema() ErrorResponse {
	log.Default().Println("SCHEMA ERROR: Something went wrong")
	return ErrorResponse{
		Message: "Something went wrong",
	}
}

func CheckSchemaErrors(result error, w http.ResponseWriter, requestBody interface{}) bool {
	if CheckJsonErrors(result, w) {
		return true
	}
	if CheckValidationErrors(w, requestBody) {
		return true
	}
	return false
}

func CheckValidationErrors(w http.ResponseWriter, requestBody interface{}) bool {
	var validate *validator.Validate
	var validationErrors validator.ValidationErrors
	validate = validator.New()
	err := validate.Struct(requestBody)
	if err != nil {
		if errors.As(err, &validationErrors) {
			erro := validationErrors[0]
			msg := ""
			switch erro.Tag() {
				case "required":
					msg = "is required"
				case "email":
					msg = "is not a valid email"
				case "min":
					msg = "must have at least " + erro.Param() + " characters"
				case "max":
					msg = "must have at most " + erro.Param() + " characters"
				default:
					msg = "is not valid"
			}
			msg = erro.Field() + " " + msg
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(FieldValidationError(msg))
			return true
		} else {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(SomethingWentWrongSchema())
			return true
		}
	}
	return false
}


func CheckJsonErrors(result error, w http.ResponseWriter) bool {
	if result != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		if errors.As(result, &syntaxError) {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(InvalidRequestBody(nil))
			return true
		} else if errors.As(result, &unmarshalTypeError) {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(InvalidField(unmarshalTypeError.Field))
			return true
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(SomethingWentWrongSchema())
		return true
	}
	return false
}