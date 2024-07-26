package database
// Getting the User profile from his id

func (db *appdbimpl) GetUserProfile(user_id int) (Profile, error) {
	var userProfile Profile
	var user User
	var followers int
	var following int

	// Fetch user profile using user_id
	user, err := db.GetUserByID(user_id)
	if err != nil {
		return userProfile, err
	}

	// Fetch the number of followers
	getFollowersQuery := `SELECT COUNT(*) FROM Follow WHERE following_id = ?`
	err = db.c.QueryRow(getFollowersQuery, user_id).Scan(&followers)
	if err != nil {
		return userProfile, err
	}

	// Fetch the number of following
	getFollowingQuery := `SELECT COUNT(*) FROM Follow WHERE followed_id = ?`
	err = db.c.QueryRow(getFollowingQuery, user_id).Scan(&following)
	if err != nil {
		return userProfile, err
	}

	// Fetch photos
	var photos []Photo
	getPhotosQuery := `SELECT photo_id, user_id, photo, caption, timestamp FROM Photo WHERE user_id = ?`
	rows, err := db.c.Query(getPhotosQuery, user_id)
	if err != nil {
		return userProfile, err
	}
	defer rows.Close()

	for rows.Next() {
		var photo Photo
		err := rows.Scan(&photo.PhotoID, &photo.UserID, &photo.Path, &photo.Caption, &photo.Timestamp)
		if err != nil {
			return userProfile, err
		}
		photos = append(photos, photo)
	}

	if err = rows.Err(); err != nil {
		return userProfile, err
	}

	// Populate profile struct
	userProfile.User = user
	userProfile.Photos = photos
	userProfile.Followers = followers
	userProfile.Following = following

	return userProfile, nil
}


func (db *appdbimpl) GetUserProfileByUsername(username string) (Profile, error) {
	var userProfile Profile
	var user User
	var followers int
	var following int

	// Get user ID from username
	var userID int
	getUserIDQuery := `SELECT user_id FROM User WHERE username = ?`
	err := db.c.QueryRow(getUserIDQuery, username).Scan(&userID)
	if err != nil {
		return userProfile, err
	}

	// Get user profile details using the user ID
	user, err = db.GetUserByID(userID)
	if err != nil {
		return userProfile, err
	}

	// Get photos related to the user
	var photos []Photo
	getPhotosQuery := `SELECT photo_id, user_id, photo, caption, timestamp FROM Photo WHERE user_id = ?`
	rows, err := db.c.Query(getPhotosQuery, userID)
	if err != nil {
		return userProfile, err
	}
	defer rows.Close()

	for rows.Next() {
		var photo Photo
		err := rows.Scan(&photo.PhotoID, &photo.UserID, &photo.Path, &photo.Caption, &photo.Timestamp)
		if err != nil {
			return userProfile, err
		}
		photos = append(photos, photo)
	}

	if err = rows.Err(); err != nil {
		return userProfile, err
	}

	// Get follower and following counts
	getFollowersQuery := `SELECT COUNT(*) FROM Follow WHERE following_id = ?`
	err = db.c.QueryRow(getFollowersQuery, userID).Scan(&followers)
	if err != nil {
		return userProfile, err
	}

	getFollowingQuery := `SELECT COUNT(*) FROM Follow WHERE followed_id = ?`
	err = db.c.QueryRow(getFollowingQuery, userID).Scan(&following)
	if err != nil {
		return userProfile, err
	}

	// Populate the Profile struct
	userProfile.User = user
	userProfile.Photos = photos
	userProfile.Followers = followers
	userProfile.Following = following

	return userProfile, nil
}

