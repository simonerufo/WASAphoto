package database
import (
	"fmt"
	"database/sql"
)

func (db *appdbimpl) DeletePhoto(user_id int, photo_id int) error {
	DELETE := `DELETE FROM Photo WHERE user_id = ? AND photo_id = ?`

	_, err := db.c.Exec(DELETE, user_id, photo_id)
	if err != nil {
		fmt.Printf("Error while trying to delete photo")
		return err
	}

	return err
}

// GetPhotoIDFromComment retrieves the photo ID associated with a specific comment ID and user ID
func (db *appdbimpl) GetPhotoIDFromComment(commentID int, userID int) (int, error) {
	QUERY := `
		SELECT c.post_id AS photo_id
		FROM Comment c
		JOIN Photo p
		ON c.post_id = p.photo_id AND c.owner_id = p.user_id
		WHERE c.comment_id = ? AND c.user_id = ?
	`

	var photoID int
	err := db.c.QueryRow(QUERY, commentID, userID).Scan(&photoID)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, fmt.Errorf("no photo found for comment ID %d and user ID %d", commentID, userID)
		}
		return 0, err
	}

	return photoID, nil
}

// CheckPhotoOwnership checks if a specific user owns the photo associated with a given comment ID
func (db *appdbimpl) CheckPhotoOwnership(commentID int, userID int) (bool, error) {
	QUERY := `
		SELECT p.user_id
		FROM Comment c
		JOIN Photo p ON c.post_id = p.photo_id
		WHERE c.comment_id = ? AND p.user_id = ?
	`

	var ownerID int
	err := db.c.QueryRow(QUERY, commentID, userID).Scan(&ownerID)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
