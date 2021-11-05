package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/orlovssky/gread/api"
	"github.com/orlovssky/gread/internal/services"
)

var authService services.AuthService

// Credentials - Holds login credentials
type Credentials struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func HandleAuthSignIn(w http.ResponseWriter, r *http.Request) {
	// Read the request
	var credentials Credentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		api.ERROR(w, http.StatusUnprocessableEntity, errors.New("could not decode JSON body"))
		return
	}

	// Check the password and get the token
	token, err := authService.SignIn(credentials.Email, credentials.Password)
	if err != nil {
		api.ERROR(w, http.StatusUnauthorized, err)
		return
	}

	api.JSON(w, http.StatusOK, map[string]interface{}{"token": token})
}
