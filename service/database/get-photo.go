package database

func (db *appdbimpl) CheckPhoto(user_id int, photo_id int) (bool, error) {
	CHECK := `SELECT COUNT(*) FROM Photo WHERE user_id = ? AND photo_id = ?`

	var photoExists bool

	err := db.c.QueryRow(CHECK, user_id, photo_id).Scan(&photoExists)
	if err != nil || !photoExists {
		return false, err
	}

	return photoExists, err

}
