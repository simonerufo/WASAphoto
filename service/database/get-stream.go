package database

import (
	"encoding/base64"
	"fmt"
)

func (db *appdbimpl) GetStream(userID int) ([]Photo, error) {
	var posts []Photo

	// SQL query to get the posts from followed users
	GETPostsFollowed := `SELECT 
							Photo.photo_id, 
							Photo.user_id, 
							Photo.photo, 
							Photo.caption, 
							Photo.timestamp
						 FROM Photo
						 JOIN Follow ON Photo.user_id = Follow.followed_id
						 WHERE Follow.following_id = ?
						 ORDER BY Photo.timestamp DESC`

	rows, err := db.c.Query(GETPostsFollowed, userID)
	if err != nil {
		fmt.Printf("error while executing query: %v\n", err)
		return posts, err
	}
	defer rows.Close()

	for rows.Next() {
		var post Photo
		var imageData []byte

		// Scan the basic post data into the Photo struct
		if err := rows.Scan(&post.PhotoID, &post.UserID, &imageData, &post.Caption, &post.Timestamp); err != nil {
			fmt.Printf("error while iterating over table: %v\n", err)
			return posts, err
		}

		// Convert image data to base64-encoded string
		post.Image = fmt.Sprintf("data:image/jpeg;base64,%s", base64.StdEncoding.EncodeToString(imageData))

		// Add the post to the slice
		posts = append(posts, post)
	}

	if err = rows.Err(); err != nil {
		fmt.Printf("Error during rows iteration: %v\n", err)
		return posts, err
	}

	return posts, nil
}
