package database

func (db *appdbimpl) GetUser(username string) (User, error) {
	var userData User
	GET := `SELECT * FROM User WHERE username= ?`

	err := db.c.QueryRow(GET, username).Scan(&userData.UserID, &userData.Username)
	if err != nil {
		return userData, err
	}

	return userData, err

}
