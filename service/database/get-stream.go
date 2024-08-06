package database

import("fmt")

func (db *appdbimpl) GetStream(user_id int) ([]Photo, error) {
	var posts []Photo
	GETPostsFollowed := `SELECT Photo.photo_id, Photo.user_id, Photo.photo, Photo.timestamp, Photo.caption
						 FROM Photo
						 JOIN Follow ON Photo.user_id = Follow.followed_id
						 WHERE Follow.following_id = ?
						 ORDER BY Photo.timestamp DESC`

	rows, err := db.c.Query(GETPostsFollowed, user_id)
	if err != nil {
		fmt.Printf("error while executing query: %v\n", err)
		return posts, err
	}
	defer rows.Close()

	for rows.Next() {
		var post Photo
		if err := rows.Scan(&post.PhotoID, &post.UserID, &post.Image, &post.Timestamp, &post.Caption); err != nil {
			fmt.Printf("error while iterating over table: %v\n", err)
			return posts, err
		}
		//post.Time = post.Timestamp.Format("2006-01-02 15:04:05")
		posts = append(posts, post)
	}

	return posts, nil
}