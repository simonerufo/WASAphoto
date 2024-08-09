package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// RemoveComment handles the deletion of a comment.
func (rt _router) RemoveComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Extract user_id from the URL parameters
	userID, err := strconv.Atoi(ps.ByName("user_id"))
	if err != nil {
		http.Error(w, "Error while fetching user ID from parameters", http.StatusBadRequest)
		return
	}

	// Extract photo_id from the URL parameters
	commentID, err := strconv.Atoi(ps.ByName("comment_id"))
	if err != nil {
		http.Error(w, "Error while fetching comment ID from parameters", http.StatusBadRequest)
		return
	}

	// Check if the user is authenticated
	Auth(w, r)

	
	// check if user owns the photo associated with the given comment ID
	isOwner,err := rt.db.CheckPhotoOwnership(commentID,userID)
	if err != nil{
		http.Error(w,"Error while checking the owner from database",http.StatusInternalServerError)
		return
	}

	if !isOwner {
		http.Error(w,"Unauthorized to delete the comment", http.StatusUnauthorized)
		return
	}

	photoID,err := rt.db.GetPhotoIDFromComment(commentID, userID)
	if err != nil{
		http.Error(w,"Error while retrieving the photo id from database", http.StatusInternalServerError)
		return
	}
	// Verify if the comment belongs to the user
	// comment, err := rt.db.GetCommentByID(comment_id)
	// if err != nil {
	// 	if err.Error() == "not found" {
	// 		http.Error(w, "Comment not found", http.StatusNotFound)
	// 	} else {
	// 		http.Error(w, "Error while retrieving comment", http.StatusInternalServerError)
	// 	}
	// 	return
	// }
	
	// user,err := rt.db.GetUserByID(userID)
	// if err != nil{
	// 	http.Error(w,"Error while retrieving user from database", http.StatusInternalServerError)
	// }

	// if comment.Username != user.Username  {
	// 	http.Error(w, "Unauthorized to delete this comment", http.StatusForbidden)
	// 	return
	// }

	// Delete the comment entry from the database
	err = rt.db.DeleteComment(userID,photoID,commentID)
	if err != nil {
		http.Error(w, "Error while deleting comment from database", http.StatusInternalServerError)
		return
	}

	// Successfully removed the comment
	fmt.Fprintln(w, "Comment successfully removed!")
}
