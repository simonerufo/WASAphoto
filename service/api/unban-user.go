package api

import (
	"fmt"
	"net/http"
	"strconv"

	"git.simonerufo.it/WASAphoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt _router) unbanUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// getto i parametri dell'user che vuole fare l'unban e quelli dell'user che viene sbannato
	user_id, err := strconv.Atoi(ps.ByName("user_id"))
	if err != nil {
		http.Error(w, "Error while parsing from parameters user id", http.StatusBadRequest)
		return
	}

	target_id, err := strconv.Atoi(ps.ByName("target_id"))
	if err != nil {
		http.Error(w, "Error while parsing from parameters target id", http.StatusBadRequest)
		return
	}

	if user_id == target_id {
		http.Error(w, "Error while trying unban yourself", http.StatusBadRequest)
		return
	}

	isBanned, err := rt.db.GetBan(user_id, target_id)
	if err != nil {
		http.Error(w, "Error while retrieving target id from db", http.StatusBadRequest)
		return
	}
	if !isBanned {
		http.Error(w, "Error while trying unban a user not banned", http.StatusBadRequest)
	}

	isAuth(w, r, user_id)

	err = rt.db.UnbanUser(user_id, target_id)
	if err != nil {
		http.Error(w, "Error while trying to remove ban relation", http.StatusBadRequest)
		return
	}

	fmt.Printf("user successfully unbanned\n")
	//non puoi sbannare te stesso
	//controlla se l'user da sbannare esiste
	// auth
	//db
	// messaggio di successo
}
