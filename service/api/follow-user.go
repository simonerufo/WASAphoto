package api

import (
	"fmt"
	"net/http"
	"strconv"

	"git.simonerufo.it/WASAphoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) followUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	//getting the http parameters
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

	//checking invalid operations
	if user_id == target_id {
		http.Error(w, "invalid action, you can't follow your self!", http.StatusBadRequest)
		return
	}

	//getting and checking if target and current users exists
	current_user, err := rt.db.GetUserByID(user_id)
	if err != nil {
		http.Error(w, "current user not recognized", http.StatusBadRequest)
		return
	}

	target_user, err := rt.db.GetUserByID(target_id)
	if err != nil {
		http.Error(w, "target user not recognized", http.StatusBadRequest)
		return
	}

	//authentication check
	isAuth(w, r, current_user.UserID)

	//adding the follow relation into database
	rt.db.FollowUser(current_user.UserID, target_user.UserID)

	//success message
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "user %s succesfully followed %s", current_user.Username, target_user.Username)
}
