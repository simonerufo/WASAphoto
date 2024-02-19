package database

// insert into Follow tables users ids of user that wants to follow another
func (db *appdbimpl) FollowUser(following_id int, followed_id int) error {
	FOLLOW := `INSERT INTO Follow(following_id,followed_id)
				VALUES(?,?)`

	_, err := db.c.Exec(FOLLOW, following_id, followed_id)
	if err != nil {
		return err
	}

	return err
}
