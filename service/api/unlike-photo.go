package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt _router) RemoveLike(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Extract user_id from the URL parameters
	userID, err := strconv.Atoi(ps.ByName("user_id"))
	if err != nil {
		http.Error(w, "Error while fetching user ID from parameters", http.StatusBadRequest)
		return
	}

	// Extract photo_id from the URL parameters
	photoID, err := strconv.Atoi(ps.ByName("photo_id"))
	if err != nil {
		http.Error(w, "Error while fetching photo ID from parameters", http.StatusBadRequest)
		return
	}

	// Check if the user is authenticated
	_,err = Auth(w,r)
	if err != nil{
		http.Error(w,"Invalid authorization",http.StatusUnauthorized)
		return
	}

	// Retrieve the owner_id from the database (needed for deleting the like)
	ownerID, err := rt.db.GetPhotoOwner(photoID)
	if err != nil {
		if err.Error() == "not found" {
			http.Error(w, "Photo not found", http.StatusNotFound)
		} else {
			http.Error(w, "Error while retrieving photo owner", http.StatusInternalServerError)
		}
		return
	}

	// Delete the like entry from the database
	err = rt.db.DeleteLike(userID, ownerID, photoID)
	if err != nil {
		http.Error(w, "Error while deleting like from database", http.StatusInternalServerError)
		return
	}

	// Successfully removed the like
	fmt.Fprintln(w, "Like successfully removed!")
}
