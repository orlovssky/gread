package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"

	"github.com/orlovssky/gread/api"
	"github.com/orlovssky/gread/internal/store"
)

// handleUserCreate - Create a user
func (s *server) handleUserCreate(w http.ResponseWriter, r *http.Request) {
	// Read the request
	var user store.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		api.ERROR(w, http.StatusUnprocessableEntity, errors.New("could not decode JSON body"))
		return
	}

	// Create user
	user, err = s.services.UserService.Create(user)
	if err != nil {
		api.ERROR(w, http.StatusConflict, err)
		return
	}

	api.JSON(w, http.StatusCreated, user)
}

// handleUserGet - Get a user
func (s *server) handleUserGet(w http.ResponseWriter, r *http.Request) {
	// Read the request
	var user store.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		api.ERROR(w, http.StatusUnprocessableEntity, errors.New("could not decode JSON body"))
		return
	}

	// Get user
	u, err := s.services.UserService.Get(user)
	if err != nil {
		api.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	api.JSON(w, http.StatusCreated, u)
}

// handleUserUpdate - Updates a user
func (s *server) handleUserUpdate(w http.ResponseWriter, r *http.Request) {
	// Parse path var
	userID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		api.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	id := r.Context().Value("id").(int)
	if id != userID {
		api.ERROR(w, http.StatusForbidden, errors.New("you do not have access"))
		return
	}
	var user store.User
	body := api.Read(w, r)
	// Read the request
	err = json.Unmarshal(body, &user)
	if err != nil {
		api.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	// Set userID that we want to update
	user.ID = userID
	fmt.Println(user)
	// // Read the request
	// user := api.Read(w, r).(store.User)
	// // Set userID that we want to update
	// user.ID = userID

	// Update user
	user, err = s.services.UserService.Update(user)
	if err != nil {
		api.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	api.JSON(w, http.StatusCreated, user)
}

// handleUserDelete - Deletes a user
func (s *server) handleUserDelete(w http.ResponseWriter, r *http.Request) {
	// Parse path var
	userID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		api.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	id := r.Context().Value("id").(int)
	if id != userID {
		api.ERROR(w, http.StatusForbidden, errors.New("you do not have access"))
		return
	}

	var user store.User
	user.ID = userID

	// Delete user
	err = s.services.UserService.Delete(user)
	if err != nil {
		api.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	api.JSON(w, http.StatusOK, "success")
}
