package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Verifying the user authorization
	_, err := Auth(w, r)
	if err != nil {
		http.Error(w, "User not authorized", http.StatusUnauthorized)
		return
	}

	// Retrieving the id of the profile that is going to be shown
	profileID, err := strconv.Atoi(ps.ByName("user_id"))
	if err != nil {
		http.Error(w, "Invalid user id", http.StatusBadRequest)
		return
	}

	// Retrieving the profile from the database
	profile, err := rt.db.GetUserProfileByID(profileID)
	if err != nil {
		http.Error(w, "Error while getting user profile", http.StatusBadGateway)
		return
	}

	// Setting the Content-Type header and encoding the profile to JSON
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(profile)
	if err != nil {
		http.Error(w, "Error while encoding user profile", http.StatusInternalServerError)
	}
}

func (rt *_router) getProfileByUsername(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Extract the username from the query parameters
	username := r.URL.Query().Get("username")
	if username == "" {
		http.Error(w, "Username is required", http.StatusBadRequest)
		return
	}

	// Fetch the user profile by username
	userProfile, err := rt.db.GetUserProfileByUsername(username)
	if err != nil {
		http.Error(w, "Error fetching user profile", http.StatusInternalServerError)
		return
	}

	// Check if the user profile is empty
	if userProfile.User.UserID == 0 {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Set response headers and encode the response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(userProfile); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}
