package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Extract the user ID and photo ID from the URL parameters
	userIDStr := ps.ByName("user_id")
	photoIDStr := ps.ByName("photo_id")

	// Convert userID and photoID to integers
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	photoID, err := strconv.Atoi(photoIDStr)
	if err != nil {
		http.Error(w, "Invalid photo ID", http.StatusBadRequest)
		return
	}

	_, err = Auth(w, r)
	if err != nil {
		http.Error(w, "Invalid authorization", http.StatusUnauthorized)
		return
	}

	// Fetch the photo details from the database
	photoData, err := rt.db.GetPhoto(userID, photoID)
	if err != nil {
		http.Error(w, "Error fetching photo", http.StatusInternalServerError)
		return
	}

	// Check if the photo is found
	if photoData == nil {
		http.Error(w, "Photo not found", http.StatusNotFound)
		return
	}

	// Convert photo data to UserPhoto struct
	photo := UserPhoto{
		PhotoID:      photoData.PhotoID,
		UserID:       photoData.UserID,
		Image:        photoData.Image,
		Timestamp:    photoData.Timestamp,
		Caption:      photoData.Caption,
		Time:         photoData.Timestamp.Format(time.RFC3339),
		LikeCount:    0,
		CommentCount: 0,
		Liked:        false,
	}

	// Set response headers and encode the response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(photo); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}

}
