package database

func (db *appdbimpl) InsertPhoto(user_id int, caption string, photo []byte) error {
	UPLOAD := `INSERT INTO Photo(user_id,caption,photo)
			VALUES(?,?,?)`

	_, err := db.c.Exec(UPLOAD, user_id, caption, photo)
	if err != nil {
		return err
	}

	return err
}
