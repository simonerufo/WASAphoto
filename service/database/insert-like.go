package database

func (db *appdbimpl) InsertLike(userID int, ownerID int, postID int) error {
	// checking if the like already exists
	var count int
	err := db.c.QueryRow("SELECT COUNT(*) FROM Like WHERE post_id = ? AND owner_id = ? AND user_id = ?", postID, ownerID, userID).Scan(&count)
	if err != nil {
		return err
	}

	if count > 0 {
		return err
	}

	// like does not already exists
	LIKE := `INSERT INTO Like (user_id, owner_id, post_id) VALUES (?, ?, ?)`

	_, err = db.c.Exec(LIKE, userID, ownerID, postID)
	if err != nil {
		return err
	}

	return nil
}
