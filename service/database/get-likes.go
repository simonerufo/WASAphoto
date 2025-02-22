package database

// GetLikesForPhoto retrieves all likes for a given photo ID
// Like represents a like on a photo
type Like struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	PhotoID  int    `json:"photo_id"`
}

// GetLikesForPhoto retrieves all likes for a given photo ID
func (db *appdbimpl) GetLikesForPhoto(photoID int) ([]Like, error) {
	query := `
        SELECT l.user_id, u.username, l.post_id AS photo_id
        FROM Like l
        JOIN User u ON l.user_id = u.user_id
        WHERE l.post_id = ?
    `

	// Execute the query
	rows, err := db.c.Query(query, photoID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Prepare a slice to hold the likes
	var likes []Like

	// Iterate through the result set
	for rows.Next() {
		var like Like
		err := rows.Scan(&like.UserID, &like.Username, &like.PhotoID)
		if err != nil {
			return nil, err
		}
		likes = append(likes, like)
	}

	// Check for any errors encountered during iteration
	if err = rows.Err(); err != nil {
		return nil, err
	}

	//Return [] if no one liked the post
	if likes == nil {
		return []Like{}, nil
	}

	return likes, nil
}
