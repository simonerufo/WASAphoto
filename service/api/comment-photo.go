package api

import (
	"fmt"
	"net/http"
	"strconv"

	"git.simonerufo.it/WASAphoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt _router) CommentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	//getto i params (user-target-postid)
	user_id, err := strconv.Atoi(ps.ByName("user_id"))
	if err != nil {
		http.Error(w, "Error while retrieving user id from params", http.StatusBadRequest)
		return
	}
	target_id, err := strconv.Atoi(ps.ByName("target_id"))
	if err != nil {
		http.Error(w, "Error while retrieving target id from params", http.StatusBadRequest)
		return
	}
	photo_id, err := strconv.Atoi(ps.ByName("photo_id"))
	if err != nil {
		http.Error(w, "Error while retrieving user id from params", http.StatusBadRequest)
		return
	}

	isAuth(w, r, user_id)
	//isAuth
	isOnline, err := rt.db.CheckPhoto(target_id, photo_id)
	if err != nil {
		http.Error(w, "Error while checking if a photo is in db", http.StatusBadRequest)
		return
	}

	if !isOnline {
		http.Error(w, "You can't comment a post that not exists", http.StatusBadRequest)
		return
	}

	err = rt.db.InsertComment(user_id, target_id, photo_id, "text")
	if err != nil {
		http.Error(w, "Error while inserting comment into db", http.StatusBadRequest)
		return
	}

	fmt.Printf("Comment successfully posted!\n")
}
