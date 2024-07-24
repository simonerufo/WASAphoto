package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	

	"git.simonerufo.it/WASAphoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)
/*
func (rt *_router) getProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	/*
		TODO:
		-DA RIVEDERE IL PATH NELL'API SWAGGER (user_id,target_id)
		- fare il parse da params dell'id
		- verificare se l'id esiste nel db
		- auth check
		-ban check(?)
	
	//getting the user id from request parameters
	//getto l'id dell'user corrente dall'header
	//getto l'id del profilo da visualizzare dai params, ne verifico la validità

	//retrieving the id of user that is searching the profile
	
	user_id, err := strconv.Atoi(r.Header.Get("Authorization"))
	if err != nil {
		http.Error(w, "Invalid Auth token", http.StatusBadGateway)
	}
	
	//auth check
	isAuth(w, r,user_id)

	fmt.Printf("uid:%d\n",user_id)
	//retrieving the id of profile that is going to be showed
	
	profileID, err := strconv.Atoi(ps.ByName("user_id"))
	if err != nil {
		http.Error(w, "invalid user id", http.StatusUnauthorized)
		return
	}
	//checking if the user searched has banned the current user
	fmt.Printf("profile_id:%d\n",profileID)
	ban, err := rt.db.GetBan(user_id, profileID)
	if err != nil {
		http.Error(w, "error checking the ban", http.StatusBadGateway)
	}
	if ban {
		fmt.Printf("user with id: %d banned user with id: %d", user_id, profileID)
		return
	}
	//mi serve il profilo preso dal db
	
	profile, err := rt.db.GetUserProfile(user_id)
	if err != nil {
		http.Error(w, "error while getting user profile", http.StatusBadGateway)
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(profile)
	if err != nil{
		http.Error(w,"error while encoding user profile",http.StatusInternalServerError)
	}
	
}
*/


func (rt *_router) getProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Retrieving the id of user that is searching the profile from the Authorization header
	authHeader := r.Header.Get("Authorization")
	fmt.Printf("Authorization Header: %s\n", authHeader)

	userID, err := strconv.Atoi(r.Header.Get("Authorization"))
	// Expecting format "Bearer <token>"
    const bearerPrefix = "Bearer "
    if !strings.HasPrefix(authHeader, bearerPrefix) {
        http.Error(w, "Invalid Authorization header format", http.StatusBadRequest)
        return
    }

    // Extract token part
    tokenStr := strings.TrimPrefix(authHeader, bearerPrefix)
    userID, err = strconv.Atoi(tokenStr)
	if err != nil {
		http.Error(w, "Invalid Auth token", http.StatusBadRequest)
		return
	}

	

	// Retrieving the id of the profile that is going to be shown
	profileID, err := strconv.Atoi(ps.ByName("user_id"))
	if err != nil {
		http.Error(w, "Invalid user id", http.StatusBadRequest)
		return
	}

	// Checking if the user searched has banned the current user
	ban, err := rt.db.GetBan(profileID, userID)
	if err != nil {
		http.Error(w, "Error checking the ban", http.StatusBadGateway)
		return
	}
	if ban {
		fmt.Printf("User with id: %d banned user with id: %d\n", profileID, userID)
		http.Error(w, "You have been banned by this user", http.StatusForbidden)
		return
	}

	// Retrieving the profile from the database
	profile, err := rt.db.GetUserProfile(profileID)
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