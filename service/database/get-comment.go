package database

func (db *appdbimpl) CheckComment(user_id int, photo_id int, comment_id int) (bool, error) {
	CHECK := ` SELECT COUNT(*) AS comment_count FROM Comment WHERE owner_id = ? 
				AND post_id = ? AND comment_id = ?`

	var commentCount int
	err := db.c.QueryRow(CHECK, user_id, photo_id, comment_id).Scan(&commentCount)
	if err != nil {
		return false, err
	}

	if commentCount == 0 {
		return false, err
	}

	return true, err
}
