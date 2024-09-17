package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	

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

/*
func (rt *_router) getProfileByUsername(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	

	userID := ps.ByName("user_id")
	if userID == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	username := r.URL.Query().Get("username")
	if username == "" {
		http.Error(w, "Username is required", http.StatusBadRequest)
		return
	}

	userProfile, err := rt.db.GetUserProfileByUsername(username)
	if err != nil {
		http.Error(w, "Error fetching user profile", http.StatusInternalServerError)
		return
	}
	if userProfile.Username == "" {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(userProfile); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
	fmt.Printf("USER: %s\n",userProfile.Username)
}
*/

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

    fmt.Printf("USER: %s\n", username)
}



