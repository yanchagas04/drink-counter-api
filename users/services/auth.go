package services

import (
	UserUtils "drink-counter-api/users/utils"
	SchemaErrors "drink-counter-api/utils/schema_errors"
	"encoding/json"
	"net/http"
	"strings"
)

type Authentication struct {
	UserId *uint
	Valid bool
}

/*
Checks if the token is valid and not expired/invalid, returning the user id if valid and an Error response if not.

# Returns:

	userId (*uint)

*/
func AuthService(w http.ResponseWriter, authHeader string, requestId uint) *uint {
	authType := strings.Split(authHeader, " ")[0]
	if authType != "Bearer" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(SchemaErrors.ErrorResponse{
			Message: "Invalid token type",
		})
		return nil
	}
	token := strings.Split(authHeader, " ")[1]
	_, err := UserUtils.ValidateToken(token)
	if err != nil {
		if strings.Contains(err.Error(), "expired") {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(SchemaErrors.ErrorResponse{
				Message: "Token expired",
			})
		} else {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(SchemaErrors.ErrorResponse{
				Message: "Invalid token",
			})
		}
		return nil
	}
	id, err := UserUtils.GetIdFromToken(token)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(SchemaErrors.ErrorResponse{
			Message: "Error getting user id from token: " + err.Error(),
		})
		return nil
	}
	if id != requestId {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(SchemaErrors.ErrorResponse{
			Message: "You don't have permission to access this resource",
		})
		return nil
	}
	return &id
}

/*
Checks if the token is valid and not expired/invalid.

# OBS:

Don't produce any Response, just the validation (true ou false).
*/
func AuthServiceValidator(w http.ResponseWriter, authHeader string, requestId uint) bool {
	authType := strings.Split(authHeader, " ")[0]
	if authType != "Bearer" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(SchemaErrors.ErrorResponse{
			Message: "Invalid token type",
		})
		return false
	}
	token := strings.Split(authHeader, " ")[1]
	_, err := UserUtils.ValidateToken(token)
	if err != nil {
		if strings.Contains(err.Error(), "expired") {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(SchemaErrors.ErrorResponse{
				Message: "Token expired",
			})
		} else {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(SchemaErrors.ErrorResponse{
				Message: "Invalid token",
			})
		}
		return false
	}
	id, err := UserUtils.GetIdFromToken(token)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(SchemaErrors.ErrorResponse{
			Message: "Error getting user id from token: " + err.Error(),
		})
		return false
	}
	if id != requestId {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(SchemaErrors.ErrorResponse{
			Message: "You don't have permission to access this resource",
		})
		return false
	}
	return true
}

/*
Checks if the user is not authorized to access private fields (ex: password). Used just for abstracting purposes.
*/
func UserNotAuthorized(w http.ResponseWriter, authHeader string, requestId uint) bool {
	return !AuthServiceValidator(w, authHeader, requestId)
}