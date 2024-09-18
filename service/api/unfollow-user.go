package api

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) unfollowUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Getting the HTTP parameters
	user_id, err := strconv.Atoi(ps.ByName("user_id"))
	if err != nil {
		http.Error(w, "Invalid user_id: cannot parse user ID", http.StatusBadRequest)
		return
	}
	following_id, err := strconv.Atoi(ps.ByName("following_id"))
	if err != nil {
		http.Error(w, "Invalid following_id: cannot parse following ID", http.StatusBadRequest)
		return
	}

	// Checking invalid operations
	if user_id == following_id {
		http.Error(w, "Invalid action: you can't unfollow yourself!", http.StatusBadRequest)
		return
	}

	// Getting and checking if current user exists
	current_user, err := rt.db.GetUserByID(user_id)
	if err != nil {
		http.Error(w, "Current user not recognized", http.StatusNotFound)
		return
	}

	// Checking if user is followed by main user
	isFollowed, err := rt.db.GetFollow(user_id, following_id)
	if err != nil {
		http.Error(w, "Error checking if the user is followed", http.StatusInternalServerError)
		return
	}
	if !isFollowed {
		http.Error(w, "Unfollow failed: user is not followed", http.StatusBadRequest)
		return
	}

	// Getting and checking if target user exists
	target_user, err := rt.db.GetUserByID(following_id)
	if err != nil {
		http.Error(w, "Target user not recognized", http.StatusNotFound)
		return
	}

	// Authentication check
	_,err = Auth(w,r)
	if err != nil{
		http.Error(w,"Invalid authorization",http.StatusUnauthorized)
		return
	}

	// Removing the follow relation from database
	err = rt.db.UnfollowUser(current_user.UserID, target_user.UserID)
	if err != nil {
		http.Error(w, "Error unfollowing user", http.StatusInternalServerError)
		return
	}

	// Success message
	w.WriteHeader(http.StatusOK)
}

