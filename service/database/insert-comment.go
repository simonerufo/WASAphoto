package database

import (
	"log"
	"time"
)

func (db *appdbimpl) InsertComment(userID int, ownerID int, postID int, text string) (int, error) {
	// Retrieve the current maximum comment_id from the Comment table
	var maxCommentID int
	err := db.c.QueryRow("SELECT COALESCE(MAX(comment_id), 0) FROM Comment").Scan(&maxCommentID)
	if err != nil {
		log.Printf("Error retrieving max comment_id: %v", err)
		return 0, err
	}

	// Increment the comment_id for the new comment
	newCommentID := maxCommentID + 1

	// Insert the new comment with the incremented comment_id
	COMMENT := `INSERT INTO Comment(comment_id, post_id, user_id, text, timestamp)
				VALUES (?, ?, ?, ?, ?)`

	_, err = db.c.Exec(COMMENT, newCommentID, postID, ownerID, text, time.Now().UTC())
	if err != nil {
		log.Printf("Error inserting comment (userID: %d, ownerID: %d, postID: %d): %v", userID, ownerID, postID, err)
		return 0, err
	}

	log.Printf("Successfully inserted comment (commentID: %d, userID: %d, postID: %d)", newCommentID, userID, postID)
	return newCommentID, nil
}
