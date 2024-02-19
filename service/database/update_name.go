package database

func (db *appdbimpl) UpdateName(user User) error {
	UPDATE := `UPDATE User SET username = ? WHERE user_id = ?`

	_, err := db.c.Exec(UPDATE, user.Username, user.UserID)
	if err != nil {
		return err
	}

	return err
}
