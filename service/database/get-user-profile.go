package database

import (
	"database/sql"
	"encoding/base64"
	"log"
)

func (db *appdbimpl) GetUserProfileByUsername(username string) (Profile, error) {
	var userProfile Profile
	var user User
	var followers, following int

	// Get user ID from username
	var userID int
	getUserIDQuery := `SELECT user_id FROM User WHERE username = ?`
	err := db.c.QueryRow(getUserIDQuery, username).Scan(&userID)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("User not found: %s", username)
			return userProfile, err
		}
		log.Printf("Error getting user ID for username %s: %v", username, err)
		return userProfile, err
	}

	// Get user profile details using the user ID
	user, err = db.GetUserByID(userID)
	if err != nil {
		log.Printf("Error fetching user profile for userID %d: %v", userID, err)
		return userProfile, err
	}

	// Get photos related to the user
	var photos []Photo
	getPhotosQuery := `SELECT photo_id, user_id, photo, caption, timestamp FROM Photo WHERE user_id = ?`
	rows, err := db.c.Query(getPhotosQuery, userID)
	if err != nil {
		log.Printf("Error retrieving photos for userID %d: %v", userID, err)
		return userProfile, err
	}
	defer rows.Close()

	for rows.Next() {
		var photo Photo
		var imageData []byte

		if err := rows.Scan(&photo.PhotoID, &photo.UserID, &imageData, &photo.Caption, &photo.Timestamp); err != nil {
			log.Printf("Error scanning photo row for userID %d: %v", userID, err)
			return userProfile, err
		}

		photo.Image = "data:image/jpeg;base64," + base64.StdEncoding.EncodeToString(imageData)
		photos = append(photos, photo)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error during rows iteration for userID %d: %v", userID, err)
		return userProfile, err
	}

	// Get follower and following counts
	getFollowersQuery := `SELECT COUNT(*) FROM Follow WHERE followed_id = ?`
	if err = db.c.QueryRow(getFollowersQuery, userID).Scan(&followers); err != nil {
		log.Printf("Error retrieving followers count for userID %d: %v", userID, err)
		return userProfile, err
	}

	getFollowingQuery := `SELECT COUNT(*) FROM Follow WHERE following_id = ?`
	if err = db.c.QueryRow(getFollowingQuery, userID).Scan(&following); err != nil {
		log.Printf("Error retrieving following count for userID %d: %v", userID, err)
		return userProfile, err
	}

	userProfile.User = user
	userProfile.Photos = photos
	userProfile.Followers = followers
	userProfile.Following = following

	return userProfile, nil
}

func (db *appdbimpl) GetUserProfileByID(profileID int) (Profile, error) {
	var userProfile Profile
	var user User
	var followers, following int

	// Fetch user details
	getUserQuery := `SELECT user_id, username FROM User WHERE user_id = ?`
	err := db.c.QueryRow(getUserQuery, profileID).Scan(&user.UserID, &user.Username)
	if err != nil {
		log.Printf("Error fetching user details for profileID %d: %v", profileID, err)
		return userProfile, err
	}

	// Fetch follower count
	getFollowersQuery := `SELECT COUNT(*) FROM Follow WHERE followed_id = ?`
	if err = db.c.QueryRow(getFollowersQuery, profileID).Scan(&followers); err != nil {
		log.Printf("Error fetching followers count for profileID %d: %v", profileID, err)
		return userProfile, err
	}

	// Fetch following count
	getFollowingQuery := `SELECT COUNT(*) FROM Follow WHERE following_id = ?`
	if err = db.c.QueryRow(getFollowingQuery, profileID).Scan(&following); err != nil {
		log.Printf("Error fetching following count for profileID %d: %v", profileID, err)
		return userProfile, err
	}

	// Fetch photos
	var photos []Photo
	getPhotosQuery := `SELECT photo_id, user_id, photo, caption, timestamp, 
                              (SELECT COUNT(*) FROM Like WHERE post_id = Photo.photo_id) AS likeCount,
                              (SELECT COUNT(*) FROM Comment WHERE post_id = Photo.photo_id) AS commentCount,
                              EXISTS(SELECT 1 FROM Like WHERE post_id = Photo.photo_id AND user_id = ?) AS liked
                       FROM Photo WHERE user_id = ?`
	rows, err := db.c.Query(getPhotosQuery, profileID, profileID)
	if err != nil {
		log.Printf("Error fetching photos for profileID %d: %v", profileID, err)
		return userProfile, err
	}
	defer rows.Close()

	for rows.Next() {
		var photo Photo
		var imageData []byte

		if err := rows.Scan(&photo.PhotoID, &photo.UserID, &imageData, &photo.Caption, &photo.Timestamp, &photo.LikeCount, &photo.CommentCount, &photo.Liked); err != nil {
			log.Printf("Error scanning photo row for profileID %d: %v", profileID, err)
			return userProfile, err
		}

		photo.Image = "data:image/jpeg;base64," + base64.StdEncoding.EncodeToString(imageData)
		photos = append(photos, photo)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error during rows iteration for profileID %d: %v", profileID, err)
		return userProfile, err
	}

	userProfile.User = user
	userProfile.Photos = photos
	userProfile.Followers = followers
	userProfile.Following = following

	return userProfile, nil
}
