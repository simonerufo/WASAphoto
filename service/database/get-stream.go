package database

import (
	"encoding/base64"
	"log"
)

func (db *appdbimpl) GetStream(userID int) ([]Photo, error) {
	var posts []Photo

	// query to get the posts from followed users
	GETPostsFollowed := `SELECT 
							Photo.photo_id, 
							Photo.user_id, 
							Photo.photo, 
							Photo.caption, 
							Photo.timestamp
					 FROM Photo
					 JOIN Follow ON Follow.followed_id = Photo.user_id
					 LEFT JOIN Ban ON Ban.banning_id = Photo.user_id AND Ban.banned_id = ?
					 WHERE Follow.following_id = ?
					 AND Ban.banning_id IS NULL
					 ORDER BY Photo.timestamp DESC`

	rows, err := db.c.Query(GETPostsFollowed, userID, userID)
	if err != nil {
		log.Printf("Error while executing query for userID %d: %v", userID, err)
		return posts, err
	}
	defer rows.Close()

	for rows.Next() {
		var post Photo
		var imageData []byte

		// Scan the basic post data into the Photo struct
		if err := rows.Scan(&post.PhotoID, &post.UserID, &imageData, &post.Caption, &post.Timestamp); err != nil {
			log.Printf("Error while scanning row for userID %d: %v", userID, err)
			return posts, err
		}

		// Convert image data to base64-encoded string
		post.Image = "data:image/jpeg;base64," + base64.StdEncoding.EncodeToString(imageData)

		// Add the post to the slice
		posts = append(posts, post)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error during rows iteration for userID %d: %v", userID, err)
		return posts, err
	}

	return posts, nil
}
