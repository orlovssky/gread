package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/orlovssky/gread/api"
	"github.com/orlovssky/gread/internal/secrets"
	"github.com/orlovssky/gread/internal/services"
	"github.com/orlovssky/gread/internal/store"
	"github.com/orlovssky/gread/pkg/auth"
)

var authService services.AuthService

// Credentials - Holds login credentials
type Credentials struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func HandleAuthSignUp(w http.ResponseWriter, r *http.Request) {
	// Read the request
	var user store.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		api.ERROR(w, http.StatusUnprocessableEntity, errors.New("could not decode JSON body"))
		return
	}

	// Create user
	user, err = userService.Create(user)
	if err != nil {
		api.ERROR(w, http.StatusConflict, err)
		return
	}

	// Get auth token
	token, err := auth.CreateToken(user.ID, secrets.LoadedSecrets.JWTSecret)
	if err != nil {
		api.ERROR(w, http.StatusConflict, err)
		return
	}

	api.JSON(w, http.StatusCreated, map[string]interface{}{"token": token})
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
