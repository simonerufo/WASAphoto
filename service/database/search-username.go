package database

func (db *appdbimpl) SearchUsername(username string) (bool, error) {
	//getting from User table the user with the username received in input
	SEARCH := `SELECT username FROM User WHERE username= ?`
	var validUser string
	//checking if username is in User table
	err := db.c.QueryRow(SEARCH, username).Scan(&validUser)
	if err != nil {
		return false, nil
	}

	return validUser != "", err
}
