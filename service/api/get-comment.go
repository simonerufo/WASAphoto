package api

import(
	"encoding/json"
	"net/http"
	"strconv"

	"git.simonerufo.it/WASAphoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter")

func (rt _router) GetPhotoComments(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Retrieve parameters from the URL
	userID, err := strconv.Atoi(ps.ByName("user_id"))
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	photoID, err := strconv.Atoi(ps.ByName("photo_id"))
	if err != nil {
		http.Error(w, "Invalid photo ID", http.StatusBadRequest)
		return
	}

	// Check if the photo exists
	isOnline, err := rt.db.CheckPhoto(userID, photoID)
	if err != nil {
		http.Error(w, "Error while checking if photo exists", http.StatusInternalServerError)
		return
	}

	if !isOnline {
		http.Error(w, "Photo not found", http.StatusNotFound)
		return
	}

	// Retrieve comments from the database
	comments, err := rt.db.GetCommentsByPhotoID(photoID)
	if err != nil {
		http.Error(w, "Error while retrieving comments", http.StatusInternalServerError)
		return
	}

	if len(comments) == 0 {
		http.Error(w, "No comments found for this photo.", http.StatusNotFound)
		return
	}

	// Send the comments in the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(comments)
}