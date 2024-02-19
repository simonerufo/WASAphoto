package database

// this function adds a new user to database and returns an User Obj
func (db *appdbimpl) InsertUser(username string) error {
	//sql query to insert into db user data
	INSERT := `INSERT INTO User(username)
			   VALUES (?);`

	//insert user into db
	_, err := db.c.Exec(INSERT, username)
	if err != nil {
		return err
	}

	return err
}
