package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) GetFollowing(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    // Retrieve user ID from parameters
    userID, err := strconv.Atoi(ps.ByName("user_id"))
    if err != nil {
        http.Error(w, "Error while fetching user ID from parameters", http.StatusBadRequest)
        return
    }

    // Check authorization
    _, err = Auth(w, r)
    if err != nil {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

    // Retrieve the list of users that the specified user follows
    following, err := rt.db.GetFollowingForUser(userID)
    if err != nil {
        if strings.Contains(err.Error(), "does not exist") {
            http.Error(w, "User not found", http.StatusNotFound)
        } else {
            http.Error(w, "Error while retrieving following users for user", http.StatusInternalServerError)
        }
        return
    }

    // Set response headers and write JSON response
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(following)
}