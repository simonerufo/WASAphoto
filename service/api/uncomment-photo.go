package api

import (
	"fmt"
	"net/http"
	"strconv"

	"git.simonerufo.it/WASAphoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt _router) UncommentPost(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	//params
	//auth
	//controllo se esiste il post
	//controllo se il commento è il mio
	//controllo se il commento esiste
	//aggiorno db
	//success
	user_id, err := strconv.Atoi(ps.ByName("user_id"))
	if err != nil {
		http.Error(w, "Error while getting target id from params", http.StatusBadRequest)
		return
	}

	target_id, err := strconv.Atoi(ps.ByName("target_id"))
	if err != nil {
		http.Error(w, "Error while getting target id from params", http.StatusBadRequest)
		return
	}

	photo_id, err := strconv.Atoi(ps.ByName("target_photoid"))
	if err != nil {
		http.Error(w, "Error while getting photo id from params", http.StatusBadRequest)
		return
	}

	comment_id, err := strconv.Atoi(ps.ByName("comment_id"))
	if err != nil {
		http.Error(w, "Error while getting comment id from params", http.StatusBadRequest)
		return
	}

	isAuth(w, r, user_id)

	isOnline, err := rt.db.CheckPhoto(target_id, photo_id)
	if err != nil {
		http.Error(w, "Error while checking if the posts exists", http.StatusInternalServerError)
		return
	}

	if !isOnline {
		http.Error(w, "Error while trying to uncomment a photo that doesn't exists", http.StatusBadRequest)
		return
	}
	//per controllare se il commento è mio devo controllare che l'user id del commento == user_id

	isPosted, err := rt.db.CheckComment(target_id, photo_id, comment_id)
	if err != nil {
		http.Error(w, "Error while checking if the comment exists", http.StatusInternalServerError)
		return
	}

	if !isPosted {
		http.Error(w, "Error while trying to delete a comment that doesn't exists", http.StatusBadRequest)
		return
	}

	err = rt.db.DeleteComment(target_id, photo_id, comment_id)
	if err != nil {
		http.Error(w, "Error while trying to delete a comment", http.StatusInternalServerError)
		return
	}

	fmt.Printf("Comment successfully deleted!\n")
}
