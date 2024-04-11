package database

func (db *appdbimpl) GetFollow(following_id int, followed_id int) (bool, error) {
	var isFollowed int
	err := db.c.QueryRow("SELECT EXISTS (SELECT 1 FROM Follow WHERE following_id = ? AND followed_id = ?)",
		following_id, followed_id).Scan(&isFollowed)
	if err != nil {
		return false, err
	}

	if isFollowed == 0 {
		return false, err
	}
	return true, err

}
