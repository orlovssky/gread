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
	// // Parse path var
	// userId, err := strconv.Atoi(chi.URLParam(r, "id"))
	// if err != nil {
	// 	api.ERROR(w, http.StatusUnprocessableEntity, err)
	// 	return
	// }

	// id := r.Context().Value("id").(int)
	// if id != userId {
	// 	api.ERROR(w, http.StatusForbidden, errors.New("you do not have access"))
	// 	return
	// }

	// // Get user
	// u, err := userService.GetById(userId)
	// if err != nil {
	// 	api.ERROR(w, http.StatusInternalServerError, err)
	// 	return
	// }

	// api.JSON(w, http.StatusOK, u)

	// =============================
	userId := r.Context().Value("id").(int)

	// Get user
	u, err := userService.GetById(userId)
	if err != nil {
		api.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	api.JSON(w, http.StatusOK, u)
}

// HandleUserUpdate - Updates a user
func HandleUserUpdate(w http.ResponseWriter, r *http.Request) {
	// Parse path var
	userId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		api.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	id := r.Context().Value("id").(int)
	if id != userId {
		api.ERROR(w, http.StatusForbidden, errors.New("you do not have access"))
		return
	}

	body := api.Read(w, r)

	// Update user
	err = userService.Update(body, userId)
	if err != nil {
		api.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	api.JSON(w, http.StatusOK, "success")
}

// HandleUserPasswordUpdate - Updates a user's password
func HandleUserPasswordUpdate(w http.ResponseWriter, r *http.Request) {
	// Parse path var
	userId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		api.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	id := r.Context().Value("id").(int)
	if id != userId {
		api.ERROR(w, http.StatusForbidden, errors.New("you do not have access"))
		return
	}

	body := api.Read(w, r)
	if body == nil {
		return
	}

	// Update user
	err = userService.UpdatePassword(body, userId)
	if err != nil {
		api.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	api.JSON(w, http.StatusOK, "success")
}

// HandleUserDelete - Deletes a user
func HandleUserDelete(w http.ResponseWriter, r *http.Request) {
	// Parse path var
	userId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		api.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	id := r.Context().Value("id").(int)
	if id != userId {
		api.ERROR(w, http.StatusForbidden, errors.New("you do not have access"))
		return
	}

	var user store.User
	user.ID = userId

	// Delete user
	err = userService.Delete(user)
	if err != nil {
		api.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	api.JSON(w, http.StatusOK, "success")
}
