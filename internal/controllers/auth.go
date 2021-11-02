package controllers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/orlovssky/gread/api"
	"github.com/orlovssky/gread/internal/services"
)

// Credentials - Holds login credentials
type Credentials struct {
	Login    string `json:"login" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// handleAuthLogin - This function checks the users email and password
// match. If they matcn an oauth token is returned
func Login(w http.ResponseWriter, r *http.Request) {
	// Read the request
	var credentials Credentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		api.ERROR(w, http.StatusUnprocessableEntity, errors.New("could not decode JSON body"))
		return
	}
	// Check the password and get the token
	token, err := services.AuthServiceInstance.Login(credentials.Login, credentials.Password)
	if err != nil {
		api.ERROR(w, http.StatusUnauthorized, err)
		return
	}

	api.JSON(w, http.StatusOK, map[string]interface{}{"token": token})
}
