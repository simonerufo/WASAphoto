package api

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) deletePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Retrieve user ID from parameters
	userID, err := strconv.Atoi(ps.ByName("user_id"))
	if err != nil {
		http.Error(w, "Error while retrieving user ID from parameters", http.StatusBadRequest)
		return
	}

	// Retrieve photo ID from parameters
	photoID, err := strconv.Atoi(ps.ByName("photo_id"))
	if err != nil {
		http.Error(w, "Error while retrieving photo ID from parameters", http.StatusBadRequest)
		return
	}

	// Auth check
	id, _ := Auth(w, r)
	if userID != id {
		http.Error(w, "User is not authorized to delete this photo", http.StatusUnauthorized)
		return
	}
	// Delete the photo from the database
	err = rt.db.DeletePhoto(userID, photoID)
	if err != nil {
		http.Error(w, "Error while trying to delete photo from database", http.StatusInternalServerError)
		return
	}

	// Set response headers and status
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
