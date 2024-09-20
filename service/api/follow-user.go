package api

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) followUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Getting the HTTP parameters
	user_id, err := strconv.Atoi(ps.ByName("user_id"))
	if err != nil {
		http.Error(w, "cannot parse user_id", http.StatusBadRequest)
		return
	}
	following_id, err := strconv.Atoi(ps.ByName("following_id"))
	if err != nil {
		http.Error(w, "cannot parse following_id", http.StatusBadRequest)
		return
	}

	// Checking invalid operations
	if user_id == following_id {
		http.Error(w, "invalid action, you can't follow yourself!", http.StatusBadRequest)
		return
	}

	// Getting and checking if target and current users exist
	current_user, err := rt.db.GetUserByID(user_id)
	if err != nil {
		http.Error(w, "current user not recognized", http.StatusBadRequest)
		return
	}

	target_user, err := rt.db.GetUserByID(following_id)
	if err != nil {
		http.Error(w, "target user not recognized", http.StatusBadRequest)
		return
	}

	// Authentication check
	_, err = Auth(w, r)
	if err != nil {
		http.Error(w, "Invalid authorization", http.StatusUnauthorized)
		return
	}

	// Adding the follow relation into database
	err = rt.db.FollowUser(current_user.UserID, target_user.UserID)
	if err != nil {
		http.Error(w, "error following user", http.StatusInternalServerError)
		return
	}

	// Success message
	w.WriteHeader(http.StatusOK)
}
