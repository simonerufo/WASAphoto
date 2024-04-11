package database

func (db *appdbimpl) DeleteComment(user_id int, photo_id int, comment_id int) error {
	UNCOMMENT := `DELETE FROM Comment WHERE user_id = ? AND post_id = ? AND comment_id = ?`

	_, err := db.c.Exec(UNCOMMENT, user_id, photo_id, comment_id)
	if err != nil {
		return err
	}
	return err
}
