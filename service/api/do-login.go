package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var userData User

	// Decode user data from body request
	err := json.NewDecoder(r.Body).Decode(&userData)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}

	// Validate username
	if !validUsername(userData.Username) {
		http.Error(w, "Invalid username", http.StatusBadRequest)
		return
	}

	// Check if the user exists in the database
	userExists := rt.db.SearchUsername(userData.Username)

	var token int
	if userExists {
		// User exists, retrieve it
		userDB, err := rt.db.GetUserByName(userData.Username)
		if err != nil {
			http.Error(w, "Error retrieving user from database", http.StatusInternalServerError)
			return
		}
		token = userDB.UserID
	} else {
		// User does not exist, create a new user
		err := rt.db.InsertUser(userData.Username)
		if err != nil {
			http.Error(w, "Error creating new user", http.StatusInternalServerError)
			return
		}
		// Retrieve the newly created user
		userDB, err := rt.db.GetUserByName(userData.Username)
		if err != nil {
			http.Error(w, "Error retrieving newly created user", http.StatusInternalServerError)
			return
		}
		token = userDB.UserID
	}

	// Set the Authorization header with a bearer token
	bearerToken := "Bearer " + strconv.Itoa(token)
	w.Header().Set("Authorization", bearerToken)

	// Send JSON response to the client
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(token)
	if err != nil {
		http.Error(w, "Error encoding token to JSON", http.StatusInternalServerError)
		return
	}

}
