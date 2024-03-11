package database

// Getting the User profile from his id
func (db *appdbimpl) GetUserProfile(user_id int) (Profile, error) {
	var userProfile Profile
	var user User
	var followers int
	var following int

	user, err := db.GetUserByID(user_id)
	if err != nil {
		return userProfile, err
	}

	//TODO: GETPhotos (prendere tutti i dati da photos)

	GETFollowers := `SELECT COUNT(*) FROM Follow WHERE following_id = ?`
	err = db.c.QueryRow(GETFollowers, user_id).Scan(&followers)
	if err != nil {
		return userProfile, err
	}

	GETFollowing := `SELECT COUNT(*) FROM Follow WHERE followed_id = ?`
	err = db.c.QueryRow(GETFollowing, user_id).Scan(&following)
	if err != nil {
		return userProfile, err
	}

	userProfile.User = user
	userProfile.Followers = followers
	userProfile.Following = following

	return userProfile, err
}
