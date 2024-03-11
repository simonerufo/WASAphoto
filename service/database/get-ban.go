package database

// GetBan checks if a user (banning id) banned another user (banned id)
func (db *appdbimpl) GetBan(banning_id int, banned_id int) (bool, error) {
	var isBanned int
	err := db.c.QueryRow("SELECT EXISTS (SELECT 1 FROM Ban WHERE banning_id = ? AND banned_id = ?)",
		banning_id, banned_id).Scan(&isBanned)
	if err != nil {
		return false, err
	}

	if isBanned == 0 {
		return false, err
	}
	return true, err
}
