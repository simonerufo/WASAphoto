package database


func (db *appdbimpl) GetFollow(following_id int, followed_id int) (bool, error) {
	var isFollowed int
	err := db.c.QueryRow("SELECT EXISTS (SELECT 1 FROM Follow WHERE following_id = ? AND followed_id = ?)",
		following_id, followed_id).Scan(&isFollowed)
	if err != nil {
		return false, err
	}

	if isFollowed == 0 {
		return false, err
	}
	return true, err

}

// retrieves the list of followers for the specified user ID
func (db *appdbimpl) GetFollowersForUser(userID int) ([]User, error) {
    // Query to get the list of followers
    query := `
        SELECT u.user_id, u.username
        FROM Follow f
        JOIN User u ON f.following_id = u.user_id
        WHERE f.followed_id = ?`

    // Execute the query
    rows, err := db.c.Query(query, userID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var followers []User

    // Iterate over the result set
    for rows.Next() {
        var user User
        if err := rows.Scan(&user.UserID, &user.Username); err != nil {
            return nil, err
        }
        followers = append(followers, user)
    }

    // Check for errors from iterating over rows
    if err := rows.Err(); err != nil {
        return nil, err
    }

    return followers, nil
}
