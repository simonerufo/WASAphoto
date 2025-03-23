package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getLikes(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Retrieve photo ID from parameters
	photoID, err := strconv.Atoi(ps.ByName("photo_id"))
	if err != nil {
		http.Error(w, "Error while fetching photo ID from parameters", http.StatusBadRequest)
		return
	}

	// Retrieve the list of likes for the specified photo
	likes, err := rt.db.GetLikesForPhoto(photoID)
	if err != nil {
		if strings.Contains(err.Error(), "does not exist") {
			http.Error(w, "Photo not found", http.StatusNotFound)
		} else {
			http.Error(w, "Error while retrieving likes for photo", http.StatusInternalServerError)
		}
		return
	}

	// Set response headers and write JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(likes)
	if err != nil {
		http.Error(w, "error while encoding the json", http.StatusInternalServerError)
	}
}
