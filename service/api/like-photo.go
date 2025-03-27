package api

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) likePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// retrieve user_id and photo_id from params
	userID, err := strconv.Atoi(ps.ByName("user_id"))
	if err != nil {
		http.Error(w, "Error while fetching user ID from parameters", http.StatusBadRequest)
		return
	}

	targetPhotoID, err := strconv.Atoi(ps.ByName("photo_id"))
	if err != nil {
		http.Error(w, "Error while fetching target photo ID from parameters", http.StatusBadRequest)
		return
	}

	// auth check
	_, err = Auth(w, r)
	if err != nil {
		http.Error(w, "Invalid authorization", http.StatusUnauthorized)
		return
	}

	// retrieve from database the owner id
	ownerID, err := rt.db.GetPhotoOwner(targetPhotoID)
	if err != nil {
		if strings.Contains(err.Error(), "does not exist") {
			http.Error(w, "Photo not found", http.StatusNotFound)
		} else {
			http.Error(w, "Error while fetching photo owner", http.StatusInternalServerError)
		}
		return
	}

	// ban checking
	isBanned, err := rt.db.GetBan(ownerID, userID)
	if err != nil {
		http.Error(w, "Error while checking if user banned you", http.StatusInternalServerError)
		return
	}

	if isBanned {
		http.Error(w, "User you're trying to like banned you!", http.StatusForbidden)
		return
	}

	// insert operation into db
	err = rt.db.InsertLike(userID, ownerID, targetPhotoID)
	if err != nil {
		http.Error(w, "Error while inserting a like entry in the database", http.StatusInternalServerError)
		return
	}

	// response
	w.WriteHeader(http.StatusNoContent)
}
