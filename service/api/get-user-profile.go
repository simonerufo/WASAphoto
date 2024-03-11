package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"git.simonerufo.it/WASAphoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	/*
		TODO:
		-DA RIVEDERE IL PATH NELL'API SWAGGER (user_id,target_id)
		- fare il parse da params dell'id
		- verificare se l'id esiste nel db
		- auth check
		-ban check(?)
	*/
	//getting the user id from request parameters
	//getto l'id dell'user corrente dall'header
	//getto l'id del profilo da visualizzare dai params, ne verifico la validità

	//retrieving the id of user that is searching the profile
	current_id, err := strconv.Atoi(r.Header.Get("Authorization"))
	if err != nil {
		http.Error(w, "Invalid Auth token", http.StatusBadGateway)
	}
	//auth check
	isAuth(w, r, current_id)

	//retrieving the id of profile that is going to be showed
	user_id, err := strconv.Atoi(ps.ByName("user_id"))
	if err != nil {
		http.Error(w, "invalid user id", http.StatusBadGateway)
		return
	}
	//checking if the user searched has banned the current user
	ban, err := rt.db.GetBan(user_id, current_id)
	if err != nil {
		http.Error(w, "error checking the ban", http.StatusBadGateway)
	}
	if ban {
		fmt.Printf("user with id: %d banned user with id: %d", user_id, current_id)
		return
	}
	//mi serve il profilo preso dal db

	profile, err := rt.db.GetUserProfile(user_id)
	if err != nil {
		http.Error(w, "error while getting user profile", http.StatusBadGateway)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profile)
	//da testare
	fmt.Printf("id : %d, name: %s,followers: %d,  following: %d\n", profile.User.UserID, profile.User.Username, profile.Followers, profile.Following)
}
