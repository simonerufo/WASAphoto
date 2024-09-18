package api

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) unbanUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Get the user_id and target_uid from the request parameters
	userID, err := strconv.Atoi(ps.ByName("user_id"))
	if err != nil {
		http.Error(w, "Error while parsing from parameters user id", http.StatusBadRequest)
		return
	}

	targetID, err := strconv.Atoi(ps.ByName("target_uid"))
	if err != nil {
		http.Error(w, "Error while parsing from parameters target id", http.StatusBadRequest)
		return
	}

	// Ensure the user is not trying to unban themselves
	if userID == targetID {
		http.Error(w, "Error while trying to unban yourself", http.StatusBadRequest)
		return
	}

	// Check if the target user is banned by the user
	isBanned, err := rt.db.GetBan(userID, targetID)
	if err != nil {
		http.Error(w, "Error while retrieving ban status from db", http.StatusInternalServerError)
		return
	}

	if !isBanned {
		http.Error(w, "Error while trying to unban a user not banned", http.StatusBadRequest)
		return
	}

	// Check if the user is authorized to perform this action
	_,err = Auth(w,r)
	if err != nil{
		http.Error(w,"Invalid authorization",http.StatusUnauthorized)
		return
	}

	// Remove the ban relationship from the database
	err = rt.db.UnbanUser(userID, targetID)
	if err != nil {
		http.Error(w, "Error while trying to remove ban relation", http.StatusInternalServerError)
		return
	}

	// Return success response
	w.WriteHeader(http.StatusNoContent)
}

