package api

import (
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
		- verificare se l'user corrente è stato bannato dal target
	*/

	//getting the user id from request parameters
	user_id, err := strconv.Atoi(ps.ByName("user_id"))
	if err != nil {
		http.Error(w, "invalid user id", http.StatusBadGateway)
		return
	}

	//getting the user from database
	user, err := rt.db.GetUserByID(user_id)
	if err != nil {
		http.Error(w, "cannot get user by id", http.StatusBadGateway)
		return
	}

	//auth check
	isAuth(w, r, user_id)
}
