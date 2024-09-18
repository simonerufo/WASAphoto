package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getMyStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    // Extract user_id from request parameters
    userID, err := strconv.Atoi(ps.ByName("user_id"))
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }

    // Check authorization
    _,err = Auth(w, r) 
	if err != nil {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

    // Retrieve the stream for the user from the database
    profile, err := rt.db.GetStream(userID)
    if err != nil {
        http.Error(w, "Profile not found", http.StatusNotFound)
        return
    }

    // Prepare and send the response
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(profile); err != nil {
        http.Error(w, "Error encoding response", http.StatusInternalServerError)
    }

}

