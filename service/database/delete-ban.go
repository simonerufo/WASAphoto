package database

func (db *appdbimpl) UnbanUser(banning_id int, banned_id int) error {
	UNBAN := `DELETE FROM Ban WHERE banning_id = ? AND banned_id = ?`

	_, err := db.c.Exec(UNBAN, banning_id, banned_id)
	if err != nil {
		return err
	}

	return err
}
