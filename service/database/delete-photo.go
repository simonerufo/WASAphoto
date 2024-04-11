package database

func (db *appdbimpl) DeletePhoto(user_id int, photo_id int) error {
	DELETE := `DELETE FROM Photo WHERE user_id = ? AND photo_id = ?`

	_, err := db.c.Exec(DELETE, user_id, photo_id)
	if err != nil {
		return err
	}

	return err
}
