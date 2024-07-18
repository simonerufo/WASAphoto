package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"git.simonerufo.it/WASAphoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var userData User
	// decoding user data from body request
	err := json.NewDecoder(r.Body).Decode(&userData)
	if err != nil {
		http.Error(w, "Error decoding json", http.StatusBadRequest)
		return
	}

	/* similar to java syntax
	if !userData.isNameValid() {
		http.Error(w, "invalid username", http.StatusBadRequest)
		return
	}*/

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
	r.Header.Set("Authorization", bearer)

	//send to client the response
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(token)
	if err != nil {
		http.Error(w, "error while encoding token", http.StatusInternalServerError)
	}
}
