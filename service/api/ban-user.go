package api

import (
	"fmt"
	"net/http"
	"strconv"

	"git.simonerufo.it/WASAphoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) banUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	/*
		TODO:
			-get dei parametri della req e cast
			-controllare se l'user vuole bannarsi da solo
			-controllare se l'user che deve essere bannato esiste
			-auth
			-add to db
			-messaggio di successo

	*/
	user_id, err := strconv.Atoi(ps.ByName("user_id"))
	if err != nil {
		http.Error(w, "cannot parse current uid", http.StatusBadRequest)
		return
	}
	target_id, err := strconv.Atoi(ps.ByName("target_id"))
	if err != nil {
		http.Error(w, "cannot parse target uid", http.StatusBadRequest)
		return
	}

	if user_id == target_id {
		http.Error(w, "you can't ban yourself!", http.StatusBadRequest)
		return
	}

	banned_user, err := rt.db.GetUserByID(target_id)
	if err != nil {
		http.Error(w, "user you're trying to ban doesn't exits", http.StatusBadRequest)
		return
	}

	isAuth(w, r, user_id)

	err = rt.db.BanUser(user_id, banned_user.UserID)
	if err != nil {
		http.Error(w, "error while banning user", http.StatusBadRequest)
	}

	fmt.Printf("user: %s successfully banned", banned_user.Username)
}
