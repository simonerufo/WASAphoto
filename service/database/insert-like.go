package database

func (db *appdbimpl) InsertLike(user_id int, owner_id int, post_id int) error {
	LIKE := `INSERT INTO Like (user_id, owner_id, post_id) VALUES (?, ?, ?) `

	_, err := db.c.Exec(LIKE, user_id, owner_id, post_id)
	if err != nil {
		return err
	}
	return err
}
