package database

func (db *appdbimpl) BanUser(banning_id int, banned_id int) error {
	BAN := `INSERT INTO Ban(banning_id,banned_id)
			VALUES(?,?)`

	_, err := db.c.Exec(BAN, banning_id, banned_id)
	if err != nil {
		return err
	}

	return err
}