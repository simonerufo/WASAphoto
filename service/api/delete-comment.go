package api

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// RemoveComment handles the deletion of a comment.
func (rt _router) uncommentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Extract user_id from the URL parameters
	userID, err := strconv.Atoi(ps.ByName("user_id"))
	if err != nil {
		http.Error(w, "Error while fetching user ID from parameters", http.StatusBadRequest)
		return
	}

	// Extract photo_id from the URL parameters
	commentID, err := strconv.Atoi(ps.ByName("photo_id"))
	if err != nil {
		http.Error(w, "Error while fetching comment ID from parameters", http.StatusBadRequest)
		return
	}

	// Check if the user is authenticated
	_, err = Auth(w, r)
	if err != nil {
		http.Error(w, "Invalid authorization", http.StatusUnauthorized)
		return
	}

	photoID, err := rt.db.GetPhotoIDFromComment(commentID, userID)
	if err != nil {
		http.Error(w, "Error while retrieving the photo id from database", http.StatusInternalServerError)
		return
	}

	// Delete the comment entry from the database
	err = rt.db.DeleteComment(userID, photoID, commentID)
	if err != nil {
		http.Error(w, "Error while deleting comment from database", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
