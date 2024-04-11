package api

import (
	"fmt"
	"net/http"
	"strconv"

	"git.simonerufo.it/WASAphoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt _router) LikePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	user_id, err := strconv.Atoi(ps.ByName("user_id"))
	if err != nil {
		http.Error(w, "Error while fetching user id from parameters", http.StatusBadRequest)
		return
	}

	target_id, err := strconv.Atoi(ps.ByName("target_id"))
	if err != nil {
		http.Error(w, "Error while fetching target id from parameters", http.StatusBadRequest)
		return
	}

	pid, err := strconv.Atoi(ps.ByName("pid"))
	if err != nil {
		http.Error(w, "Error while fetching photo id from parameters", http.StatusBadRequest)
		return
	}

	isAuth(w, r, user_id)

	isBanned, err := rt.db.GetBan(target_id, user_id)
	if err != nil {
		http.Error(w, "Error while checking if user banned you", http.StatusBadRequest)
		return
	}

	if isBanned {
		http.Error(w, "User you're trying to like banned you!", http.StatusBadRequest)
		return
	}

	postExist, err := rt.db.CheckPhoto(target_id, pid)
	if err != nil {
		http.Error(w, "Error while trying to check if post exist", http.StatusBadRequest)
		return
	}
	if !postExist {
		http.Error(w, "You can't like a post that doesn't exist!", http.StatusBadRequest)
		return
	}

	err = rt.db.InsertLike(user_id, target_id, pid)
	if err != nil {
		http.Error(w, "Error while inserting a like entry in db", http.StatusBadRequest)
		return
	}
	//success
	fmt.Printf("Post successfully liked!\n")
}
