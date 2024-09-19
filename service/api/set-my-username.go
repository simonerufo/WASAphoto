package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

/*
updates the username using a PUT request to the server
*/
func (rt *_router) setMyUserName(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	var userData User

	// getting user id that's calling the request
	uid, err := strconv.Atoi(ps.ByName("user_id"))
	if err != nil {
		http.Error(w, "invalid param name", http.StatusBadRequest)
		return
	}

	// decoding request body into userData and checking its validity
	err = json.NewDecoder(r.Body).Decode(&userData)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	// retrieving user entry from database
	userDB, err := rt.db.GetUserByID(uid)
	if err != nil {
		http.Error(w, "cannot get user from database", http.StatusBadRequest)
		return
	}

	// checking the authorization
	userID, err := Auth(w, r)
	if err != nil {
		http.Error(w, "Invalid authorization", http.StatusUnauthorized)
		return
	}

	// the user can update just his name
	if userID != uid {
		http.Error(w, "Cannot update this username", http.StatusUnauthorized)
		return
	}

	// checking if new username is valid or no
	if !validUsername(userData.Username) {
		http.Error(w, "invalid username", http.StatusBadRequest)
		return
	}

	// checking if new username was already in db
	taken := rt.db.SearchUsername(userData.Username)
	fmt.Printf("uD: %s, uDB: %s \n", userData.Username, userDB.Username)
	if taken {
		if userDB.Username != userData.Username {
			http.Error(w, "name already taken", http.StatusBadRequest)
			return
		}
	}

	// updating username in database
	userDB.Username = userData.Username
	err = rt.db.UpdateName(userDB)
	if err != nil {
		http.Error(w, "bad request by updating username into database", http.StatusBadRequest)
		return
	}

	// printing the response
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "username succesfully updated to: %s\n", userData.Username)
}
