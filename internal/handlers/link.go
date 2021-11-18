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

var linkService services.LinkService

// HandleLinkPost - Parse a link
func HandleLinkPost(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("id").(int)

	var ltp store.LinkToParse
	err := json.NewDecoder(r.Body).Decode(&ltp)
	if err != nil {
		api.ERROR(w, http.StatusUnprocessableEntity, errors.New("could not decode JSON body"))
		return
	}

	// Parse link
	link, err := linkService.Create(ltp, userId)
	if err != nil {
		api.ERROR(w, http.StatusConflict, err)
		return
	}

	api.JSON(w, http.StatusCreated, link)
}

// HandleLinksGet - Get links
func HandleLinksGet(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("id").(int)

	// Get links
	links, err := linkService.GetList(userId)
	if err != nil {
		api.ERROR(w, http.StatusConflict, err)
		return
	}

	api.JSON(w, http.StatusOK, links)
}

// HandleLinkDelete - Delete link
func HandleLinkDelete(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("id").(int)

	// Parse path var
	linkId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		api.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	// Delete user
	err = linkService.Delete(linkId, userId)
	if err != nil {
		api.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	api.JSON(w, http.StatusOK, "success")
}
