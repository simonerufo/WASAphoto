package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

func (rt _router) commentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Get the params (user_id and photo_id)
	userID, err := strconv.Atoi(ps.ByName("user_id"))
	if err != nil {
		http.Error(w, "Error while retrieving user id from params", http.StatusBadRequest)
		return
	}

	photoID, err := strconv.Atoi(ps.ByName("photo_id"))
	if err != nil {
		http.Error(w, "Error while retrieving photo id from params", http.StatusBadRequest)
		return
	}

	// Authenticate the user
	_, err = Auth(w, r)
	if err != nil {
		http.Error(w, "Invalid authorization", http.StatusUnauthorized)
		return
	}

	// Parse the comment text from the request body
	var commentReq CommentRequest
	err = json.NewDecoder(r.Body).Decode(&commentReq)
	if err != nil || commentReq.CommentText == "" {
		http.Error(w, "Invalid comment data", http.StatusBadRequest)
		return
	}

	// Insert the comment into the database
	commentID, err := rt.db.InsertComment(userID, userID, photoID, commentReq.CommentText)
	if err != nil {
		http.Error(w, "Error while inserting comment into db", http.StatusInternalServerError)
		return
	}

	// Prepare the response
	user, err := rt.db.GetUserByID(userID)
	if err != nil {
		http.Error(w, "Error while retrieving username from db", http.StatusInternalServerError)
		return
	}

	commentRes := Comment{
		Username:    user.Username,
		CommentID:   commentID,
		CommentText: commentReq.CommentText,
		Timestamp:   time.Now().UTC().Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(commentRes)
	if err != nil {
		http.Error(w, "error while encoding the json", http.StatusBadRequest)
	}

}
