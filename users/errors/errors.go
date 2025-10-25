package errors

import (
	"drink-counter-api/utils"
)

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