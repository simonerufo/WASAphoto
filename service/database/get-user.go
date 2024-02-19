package database

func (db *appdbimpl) GetUserByName(username string) (User, error) {
	var userData User
	GET := `SELECT * FROM User WHERE username= ?`

	err := db.c.QueryRow(GET, username).Scan(&userData.UserID, &userData.Username)
	if err != nil {
		return userData, err
	}

	return userData, err

}

func (db *appdbimpl) GetUserByID(id int) (User, error) {
	var userData User
	GET := `SELECT * FROM User WHERE user_id= ?`

	err := db.c.QueryRow(GET, id).Scan(&userData.UserID, &userData.Username)
	if err != nil {
		return userData, err
	}

	return userData, err

}
