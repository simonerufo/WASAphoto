package database

import (
	"database/sql"
	"fmt"
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
			return Comment{}, fmt.Errorf("comment with ID %d not found", commentID)
		}
		return Comment{}, fmt.Errorf("error retrieving comment: %v", err)
	}

	return comment, nil
}

func (db *appdbimpl) GetCommentsByPhotoID(postID int) ([]Comment, error) {
	QUERY := `SELECT c.comment_id, u.username, c.text, c.timestamp
			  FROM Comment c
			  JOIN User u ON c.user_id = u.user_id
			  WHERE c.post_id = ?`
	rows, err := db.c.Query(QUERY, postID)
	if err != nil {
		return nil, fmt.Errorf("error querying comments: %w", err)
	}
	defer rows.Close()

	var comments []Comment
	for rows.Next() {
		var comment Comment
		err := rows.Scan(&comment.CommentID, &comment.Username, &comment.CommentText, &comment.Timestamp)
		if err != nil {
			return nil, fmt.Errorf("error scanning comment row: %w", err)
		}
		comments = append(comments, comment)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error with rows: %w", err)
	}

	//return [], if None has commented the post yet
	if len(comments) == 0 {
		return []Comment{}, nil
	}
	return comments, nil
}
