package api

import (
	"fmt"
	"net/http"
	"strconv"

	"git.simonerufo.it/WASAphoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt _router) unlikePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
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

	photo_id, err := strconv.Atoi(ps.ByName("photo_id"))
	if err != nil {
		http.Error(w, "Error while fetching photo id from parameters", http.StatusBadRequest)
		return
	}

	isAuth(w, r, user_id)

	photoExists, err := rt.db.CheckPhoto(target_id, photo_id)
	if err != nil {
		http.Error(w, "Error while checking if post exists", http.StatusBadRequest)
		return
	}

	if !photoExists {
		http.Error(w, "You can't unlike a post that doesn't even exist!", http.StatusBadRequest)
		return
	}

	err = rt.db.DeleteLike(user_id, target_id, photo_id)
	if err != nil {
		http.Error(w, "Error while deleting like from db", http.StatusBadRequest)
		return
	}

	fmt.Printf("Post successfully unliked!\n")
}
