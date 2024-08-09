package api

import (
	"fmt"
	"net/http"
	"strconv"
	"encoding/json"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) BanUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if userID == targetID {
		http.Error(w, "you can't ban yourself!", http.StatusBadRequest)
		return
	}

	bannedUser, err := rt.db.GetUserByID(targetID)
	if err != nil {
		http.Error(w, "user you're trying to ban doesn't exist", http.StatusNotFound)
		return
	}

	Auth(w,r)

	err = rt.db.BanUser(userID, bannedUser.UserID)
	if err != nil {
		http.Error(w, "error while banning user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bannedUser)

	fmt.Printf("user: %s successfully banned", bannedUser.Username)
}

