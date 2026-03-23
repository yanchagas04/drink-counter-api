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
AuthService checks if the token is valid and not expired.

Returns:

	Authentication {
		UserId: *uint (user id if valid, else nil)
		Valid: bool (true if valid, else false)
	}

If invalid, returns false and send an unathorized response.
*/
func AuthService(w http.ResponseWriter, authHeader string) Authentication {
	authType := strings.Split(authHeader, " ")[0]
	if authType != "Bearer" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(SchemaErrors.ErrorResponse{
			Message: "Invalid token type",
		})
		return Authentication{
			UserId: nil,
			Valid: false,
		}
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
		return Authentication{
			UserId: nil,
			Valid: false,
		}
	}
	id, err := UserUtils.GetIdFromToken(token)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(SchemaErrors.ErrorResponse{
			Message: "Error getting user id from token: " + err.Error(),
		})
		return Authentication{
			UserId: nil,
			Valid: false,
		}
	}
	return Authentication{
		UserId: &id,
		Valid: true,
	}
}