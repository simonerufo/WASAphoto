package database

import("fmt")

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

func (db *appdbimpl) GetBanList(userID int) ([]int, error) {
	var bannedUserIDs []int

	query := "SELECT banned_id FROM Ban WHERE banning_id = ?"

	rows, err := db.c.Query(query, userID)
	if err != nil {
		
		return nil, fmt.Errorf("error while executing the query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var bannedID int
		if err := rows.Scan(&bannedID); err != nil {
			return nil, fmt.Errorf("error while retrieving user ids from table: %w", err)
		}
		bannedUserIDs = append(bannedUserIDs, bannedID)
	}

	if err := rows.Err(); 
	err != nil {
		return nil, fmt.Errorf("error : %w", err)
	}

	return bannedUserIDs, nil
}
