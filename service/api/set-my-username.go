package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"git.simonerufo.it/WASAphoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

/*
updates the username using a PUT request to the server
*/
func (rt *_router) setMyUsername(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	/*
		TODO:
			-richiedere l'id e verificare se è lo stesso dell'utente che ha fatto la req (ps.ByName())
			-richiedere il body della richiesta per cambiare l'username (json encode)
			-verificare che il nome sia valido
			-aggiornare la entry del db dell'user
			-messaggio di risposta
	*/
	var userData User

	//getting user id that's calling the request
	uid, err := strconv.Atoi(ps.ByName("user_id"))
	if err != nil {
		http.Error(w, "invalid param name", http.StatusBadRequest)
		return
	}

	//decoding request body into userData and checking its validity
	err = json.NewDecoder(r.Body).Decode(&userData)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	//retrieving user entry from database
	fmt.Printf("uid: %d\n", uid)
	userDB, err := rt.db.GetUserByID(uid)
	if err != nil {
		http.Error(w, "cannot get user from database", http.StatusBadRequest)
		return
	}

	//checking the authorization
	Auth(w,r);
	
	//checking if new username is valid or no
	if !validUsername(userData.Username) {
		http.Error(w, "invalid username", http.StatusBadRequest)
		return
	}

	//checking if new username was already in db
	valid := rt.db.SearchUsername(userData.Username)
	if valid {
		http.Error(w, "name already taken", http.StatusBadRequest)
		return
	}

	//updating username in database
	userDB.Username = userData.Username
	err = rt.db.UpdateName(userDB)
	if err != nil {
		http.Error(w, "bad request by updating username into database", http.StatusBadRequest)
		return
	}

	//printing the response
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "username succesfully updated to: %s\n", userData.Username)
}
