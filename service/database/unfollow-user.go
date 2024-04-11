package database

func (db *appdbimpl) UnfollowUser(following_id int, followed_id int) error {
	UNFOLLOW := `DELETE FROM Follow WHERE following_id = ? AND followed_id = ?`

	_, err := db.c.Exec(UNFOLLOW, following_id, followed_id)
	if err != nil {
		return err
	}

	return err
}
