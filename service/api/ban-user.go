package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) banUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// parsing user id and target user id from params
	userID, err := strconv.Atoi(ps.ByName("user_id"))
	if err != nil {
		http.Error(w, "cannot parse user_id", http.StatusBadRequest)
		return
	}
	targetID, err := strconv.Atoi(ps.ByName("target_uid"))
	if err != nil {
		http.Error(w, "cannot parse target_uid", http.StatusBadRequest)
		return
	}

	// auto ban check
	if userID == targetID {
		http.Error(w, "you can't ban yourself!", http.StatusBadRequest)
		return
	}

	// ban a user not found
	bannedUser, err := rt.db.GetUserByID(targetID)
	if err != nil {
		http.Error(w, "user you're trying to ban doesn't exist", http.StatusNotFound)
		return
	}

	// auth check
	_, err = Auth(w, r)
	if err != nil {
		http.Error(w, "Invalid authorization", http.StatusUnauthorized)
		return
	}

	// database error
	err = rt.db.BanUser(userID, bannedUser.UserID)
	if err != nil {
		http.Error(w, "error while banning user", http.StatusInternalServerError)
		return
	}

	// ok response
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(bannedUser)
	if err != nil {
		http.Error(w, "error while encoding the json", http.StatusBadRequest)
	}
}
