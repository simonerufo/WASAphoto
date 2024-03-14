package database

func (db *appdbimpl) GetStream(user_id int) ([]Post, error) {
	var posts []Post
	GETPostsFollowed := `SELECT Photo.*
						 FROM Photo
						 JOIN Follow ON Photo.user_id = Follow.followed_id
						 WHERE Follow.following_id = ?
						 ORDER BY Photo.timestamp DESC`

	rows, err := db.c.Query(GETPostsFollowed, user_id)
	if err != nil {
		return posts, err
	}
	defer rows.Close()

	for rows.Next() {
		var post Post
		var user User
		if err := rows.Scan(&post.PostID, &user.UserID, &post.Caption, &post.Timestamp); err != nil {
			return posts, err
		}
	}

	return posts, err
}
