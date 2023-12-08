package database

// this function adds a new user to database and returns an User Obj
func (db *appdbimpl) InsertUser(user User) (User, error) {
	//sql query to insert into db user data
	INSERT := `INSERT INTO User(username)
			   VALUES (?);`

	var usr User
	usr.Username = user.Username
	//insert user into db
	_, err := db.c.Exec(INSERT, usr.Username)
	if err != nil {
		return usr, err
	}
	usr.UserID = user.UserID
	return usr, err
}
