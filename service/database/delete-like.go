package database

func (db *appdbimpl) DeleteLike(user_id int, owner_id int, post_id int) error {
	UNLIKE := `DELETE FROM Like WHERE user_id = ? AND owner_id = ? AND post_id = ?`

	_, err := db.c.Exec(UNLIKE, user_id, owner_id, post_id)
	if err != nil {
		return err
	}

	return err
}
