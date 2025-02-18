package database

import (
	"fmt"
)

func (db *appdbimpl) InsertComment(userID int, ownerID int, postID int, text string) (int, error) {
	// Retrieve the current maximum comment_id from the Comment table
	var maxCommentID int
	err := db.c.QueryRow("SELECT COALESCE(MAX(comment_id), 0) FROM Comment").Scan(&maxCommentID)
	if err != nil {
		return 0, fmt.Errorf("error retrieving max comment_id: %w", err)
	}

	// Increment the comment_id for the new comment
	newCommentID := maxCommentID + 1

	// Insert the new comment with the incremented comment_id
	COMMENT := `INSERT INTO Comment(comment_id, post_id, user_id, text, timestamp)
				VALUES (?, ?, ?, ?, ?)`

	_, err = db.c.Exec(COMMENT, newCommentID, postID, ownerID, userID, text)
	if err != nil {
		return 0, fmt.Errorf("error inserting comment: %w", err)
	}

	return newCommentID, nil
}
