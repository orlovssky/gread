package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"

	"github.com/orlovssky/gread/api"
	"github.com/orlovssky/gread/internal/services"
	"github.com/orlovssky/gread/internal/store"
)

var userService services.UserService

// HandleUserCreate - Create a user
func HandleUserCreate(w http.ResponseWriter, r *http.Request) {
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

	api.JSON(w, http.StatusCreated, user)
}

// HandleUserGet - Get a user
func HandleUserGet(w http.ResponseWriter, r *http.Request) {
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

	// Get user
	u, err := userService.GetById(userID)
	if err != nil {
		api.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	api.JSON(w, http.StatusOK, u)
}

// HandleUserUpdate - Updates a user
// func HandleUserUpdate(w http.ResponseWriter, r *http.Request) {
// 	// Parse path var
// 	userID, err := strconv.Atoi(chi.URLParam(r, "id"))
// 	if err != nil {
// 		api.ERROR(w, http.StatusUnprocessableEntity, err)
// 		return
// 	}

// 	id := r.Context().Value("id").(int)
// 	if id != userID {
// 		api.ERROR(w, http.StatusForbidden, errors.New("you do not have access"))
// 		return
// 	}

// 	body := api.Read(w, r)

// 	// Update user
// 	user, err := userService.Update(body, userID)
// 	if err != nil {
// 		api.ERROR(w, http.StatusInternalServerError, err)
// 		return
// 	}

// 	api.JSON(w, http.StatusCreated, user)
// }

// HandleUserDelete - Deletes a user
func HandleUserDelete(w http.ResponseWriter, r *http.Request) {
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
	err = userService.Delete(user)
	if err != nil {
		api.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	api.JSON(w, http.StatusOK, "success")
}