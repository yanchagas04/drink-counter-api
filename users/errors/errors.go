package errors

import (
	SchemaErrors "drink-counter-api/utils/schema_errors"
)

func UserNotFound() SchemaErrors.ErrorResponse {
	return SchemaErrors.ErrorResponse{
		Message: "User not found",
	}
}

func UserAlreadyExists() SchemaErrors.ErrorResponse {
	return SchemaErrors.ErrorResponse{
		Message: "User already exists",
	}
}