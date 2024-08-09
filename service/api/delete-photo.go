package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)
/*
func (rt *_router) deletePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// getto params (id utente id post)
	// mi autentico
	// controllo che la foto esista su db
	// elimino
	// success
	user_id, err := strconv.Atoi(ps.ByName("user_id"))
	if err != nil {
		http.Error(w, "Error while retrieving user id from parameters", http.StatusBadRequest)
		return
	}

	pid, err := strconv.Atoi(ps.ByName("pid"))
	if err != nil {
		http.Error(w, "Error while retrieving photo id from parameters", http.StatusBadRequest)
		return
	}

	isAuth(w, r, user_id)

	isUploaded, err := rt.db.CheckPhoto(user_id, pid)
	if err != nil {
		http.Error(w, "Error while retrieving photo from db", http.StatusBadRequest)
		return
	}

	if !isUploaded {
		http.Error(w, "Error while trying to delete a post that doesn't exist", http.StatusBadRequest)
		return
	}

	err = rt.db.DeletePhoto(user_id, pid)
	if err != nil {
		http.Error(w, "Error while deleting a photo from db", http.StatusBadRequest)
		return
	}

	fmt.Printf("photo successfully deleted!\n")

}
*/
func (rt *_router) deletePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    // Retrieve user ID from parameters
    userID, err := strconv.Atoi(ps.ByName("user_id"))
    if err != nil {
        http.Error(w, "Error while retrieving user ID from parameters", http.StatusBadRequest)
        return
    }

    // Retrieve photo ID from parameters
    photoID, err := strconv.Atoi(ps.ByName("photo_id"))
    if err != nil {
        http.Error(w, "Error while retrieving photo ID from parameters", http.StatusBadRequest)
        return
    }

    //Auth check
	id,_ := Auth(w,r)
	if userID != id{
		http.Error(w,"User is not authorized to delete this photo",http.StatusUnauthorized)
		return
	}
    // Delete the photo from the database
    err = rt.db.DeletePhoto(userID, photoID)
    if err != nil {
        http.Error(w, "Error while trying to delete photo from database", http.StatusInternalServerError)
        return
    }

    // Set response headers and status
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusNoContent) // No Content indicates successful deletion without returning content
	fmt.Printf("Photo successfully deleted")
}
