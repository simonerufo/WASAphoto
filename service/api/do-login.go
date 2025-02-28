package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
// 	var userData User
// 	// decoding user data from body request
// 	err := json.NewDecoder(r.Body).Decode(&userData)
// 	if err != nil {
// 		http.Error(w, "Error decoding json", http.StatusBadRequest)
// 		return
// 	}

/*
	 similar to java syntax
		if !userData.isNameValid() {
			http.Error(w, "invalid username", http.StatusBadRequest)
			return
		}

		//checking if is a valid username
		if !validUsername(userData.Username) {
			http.Error(w, "invalid username", http.StatusBadRequest)
			return
		}
		//checking if user exists (if not, create it)
		valid := rt.db.SearchUsername(userData.Username)
		if err != nil {
			http.Error(w, "error while searching username in db", http.StatusBadRequest)
			return
		}

		//getting user id as a token on user creation/retrieve
		var token int

		if valid {
			userDB, err := rt.db.GetUserByName(userData.Username)
			if err != nil {
				http.Error(w, "can't get the user", http.StatusBadRequest)
				return
			}
			//assign the token from database
			token = userDB.UserID
			//printing the response
			//w.WriteHeader(http.StatusOK)
			//fmt.Fprintf(w, "user %s successfully logged\n", userData.Username)

		} else {

			err := rt.db.InsertUser(userData.Username)
			if err != nil {
				http.Error(w, "can't create the user", http.StatusBadRequest)
				return
			}
			userDB, err := rt.db.GetUserByName(userData.Username)
			if err != nil {
				http.Error(w, "can't get the user", http.StatusBadRequest)
				return
			}
			//assign the token from database
			token = userDB.UserID
			//printing the response
			//w.WriteHeader(http.StatusOK)
			//fmt.Fprintf(w, "user %s successfully created\n", userData.Username)

		}

		bearer := strconv.Itoa(token)
		//authorizing the user
		r.Header.Set("Authorization ", bearer)

		//send to client the response
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusCreated)
		err = json.NewEncoder(w).Encode(token)
		if err != nil {
			http.Error(w, "error while encoding token", http.StatusInternalServerError)
		}
		fmt.Printf("id:%d, user:%s\n",token,userData.Username)
	}
*/
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
