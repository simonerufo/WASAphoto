package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getBanList(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Extract user_id from request parameters
	userID, err := strconv.Atoi(ps.ByName("user_id"))
	if err != nil {
		http.Error(w, "Error while parsing user id", http.StatusBadRequest)
		return
	}

	// Retrieve the ban list for the user from the database
	bannedUserIDs, err := rt.db.GetBanList(userID)
	if err != nil {
		http.Error(w, "Error while retrieving ban list from database", http.StatusInternalServerError)
		return
	}

	// Prepare the response data
	response := map[string][]int{
		"banned_user_ids": bannedUserIDs,
	}

	// Set the content type and write the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}

func (rt *_router) getBan(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userIDStr := ps.ByName("user_id")       // Extract user_id from the request
	targetUIDStr := ps.ByName("target_uid") // Extract target_uid from the request

	// Convert string IDs to integers
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user_id", http.StatusBadRequest)
		return
	}
	targetUID, err := strconv.Atoi(targetUIDStr)
	if err != nil {
		http.Error(w, "Invalid target_uid", http.StatusBadRequest)
		return
	}

	// Check if the target user is banned
	isBanned, err := rt.db.GetBan(userID, targetUID)
	if err != nil {
		http.Error(w, "Failed to check ban status", http.StatusInternalServerError)
		return
	}

	// Prepare response
	response := map[string]bool{"is_banned": isBanned}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
