package database

import (
	"log"
)

// InsertPhoto inserts a new photo into the Photo table with a manually managed photo_id and returns the photo ID.
func (db *appdbimpl) InsertPhoto(userID int, caption string, photo []byte) (int, error) {
	// Fetch the current maximum photo_id
	var maxID int
	err := db.c.QueryRow(`SELECT COALESCE(MAX(photo_id), 0) FROM Photo`).Scan(&maxID)
	if err != nil {
		log.Printf("Error fetching max photo_id: %v", err)
		return 0, err
	}

	// Increment the ID
	newPhotoID := maxID + 1

	// Insert the new photo record with the incremented photo_id
	UPLOAD := `INSERT INTO Photo(photo_id, user_id, caption, photo) 
               VALUES (?, ?, ?, ?)`

	_, err = db.c.Exec(UPLOAD, newPhotoID, userID, caption, photo)
	if err != nil {
		log.Printf("Error inserting photo (userID: %d, caption: %s): %v", userID, caption, err)
		return 0, err
	}

	log.Printf("Successfully inserted photo (photoID: %d, userID: %d)", newPhotoID, userID)
	// Return the new photo ID
	return newPhotoID, nil
}
