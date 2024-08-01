package database
import ("fmt")

func (db *appdbimpl) InsertLike(userID int, ownerID int, postID int) error {
	// First, check if the like already exists
	var count int
	err := db.c.QueryRow("SELECT COUNT(*) FROM Like WHERE post_id = ? AND owner_id = ? AND user_id = ?", postID, ownerID, userID).Scan(&count)
	if err != nil {
		fmt.Printf("error checking like existence: %w", err)
		return err
	}

	if count > 0 {
		// If like already exists, return an error or handle as needed
		fmt.Printf("like already exists for post_id=%d, owner_id=%d, user_id=%d", postID, ownerID, userID)
		return err
	}

	// Proceed with insertion if like does not already exist
	LIKE := `INSERT INTO Like (user_id, owner_id, post_id) VALUES (?, ?, ?)`

	_, err = db.c.Exec(LIKE, userID, ownerID, postID)
	if err != nil {
		fmt.Printf("error executing query: %w", err)
		return err
	}

	return nil
}

