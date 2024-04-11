package database

func (db *appdbimpl) InsertComment(user_id int, owner_id int, photo_id int, text string) error {

	COMMENT := `INSERT INTO Comment(photo_id,owner_id,user_id,text)
				VALUES(?,?,?,?)`

	_, err := db.c.Exec(COMMENT, photo_id, owner_id, user_id, text)
	if err != nil {
		return err
	}

	return err
}
