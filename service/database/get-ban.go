package database

// GetName is an example that shows you how to query data
func (db *appdbimpl) GetBan(banning_id int, banned_id int) (bool, error) {
	var isBanned int
	err := db.c.QueryRow("SELECT EXISTS (SELECT 1 FROM Ban WHERE banning_id = ? AND banned_id = ?)",
		banning_id, banned_id).Scan(&isBanned)
	if err != nil {
		return false, err
	}

	if isBanned == 1 {
		return true, err
	} else {
		return false, err
	}

}
