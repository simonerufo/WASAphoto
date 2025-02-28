package database

import (
	"database/sql"
	"errors"
	"log"
)

func (db *appdbimpl) CheckComment(user_id int, photo_id int, comment_id int) (bool, error) {
	CHECK := ` SELECT COUNT(*) AS comment_count FROM Comment WHERE owner_id = ? 
				AND post_id = ? AND comment_id = ?`

	var commentCount int
	err := db.c.QueryRow(CHECK, user_id, photo_id, comment_id).Scan(&commentCount)
	if err != nil {
		return false, err
	}

	if commentCount == 0 {
		return false, err
	}

	return true, err
}

// GetCommentByID retrieves a comment by its ID from the database
func (db *appdbimpl) GetCommentByID(commentID int) (Comment, error) {
	QUERY := `
		SELECT c.comment_id, c.text, c.timestamp, u.username
		FROM Comment c
		INNER JOIN User u ON c.user_id = u.user_id
		WHERE c.comment_id = ?`

	row := db.c.QueryRow(QUERY, commentID)

	var comment Comment

	err := row.Scan(&comment.CommentID, &comment.CommentText, &comment.Timestamp, &comment.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("comment with ID %d not found", commentID)
			return Comment{}, errors.New("comment not found")
		}
		log.Printf("error retrieving comment with ID %d: %v", commentID, err)
		return Comment{}, errors.New("failed to retrieve comment")
	}

	return comment, nil
}

func (db *appdbimpl) GetCommentsByPhotoID(postID int) ([]Comment, error) {
	QUERY := `
		SELECT c.comment_id, u.username, c.text, c.timestamp
		FROM Comment c
		JOIN User u ON c.user_id = u.user_id
		WHERE c.post_id = ?`

	rows, err := db.c.Query(QUERY, postID)
	if err != nil {
		log.Printf("error querying comments for post ID %d: %v", postID, err)
		return nil, errors.New("failed to query comments")
	}
	defer rows.Close()

	var comments []Comment
	for rows.Next() {
		var comment Comment
		err := rows.Scan(&comment.CommentID, &comment.Username, &comment.CommentText, &comment.Timestamp)
		if err != nil {
			log.Printf("error scanning comment row for post ID %d: %v", postID, err)
			return nil, errors.New("failed to scan comment row")
		}
		comments = append(comments, comment)
	}

	if err = rows.Err(); err != nil {
		log.Printf("error processing rows for post ID %d: %v", postID, err)
		return nil, errors.New("error occurred while processing comments")
	}

	// return empty slice instead of nil if there are no comments
	if len(comments) == 0 {
		return []Comment{}, nil
	}
	return comments, nil
}
