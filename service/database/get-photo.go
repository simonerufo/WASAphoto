package database

import (
	"database/sql"
	"encoding/base64"
	"fmt"
)

/*
	func (db *appdbimpl) GetPhoto(userID int, photoID int) (*Photo, error) {
		// Define the SQL query
		query := `SELECT photo_id, user_id, caption, photo, timestamp
				  FROM Photo
				  WHERE user_id = ? AND photo_id = ?`

		// Execute the query
		row := db.c.QueryRow(query, userID, photoID)

		// Create a Photo instance to hold the result
		var photo Photo

		// Scan the result into the Photo instance
		err := row.Scan(&photo.PhotoID, &photo.UserID, &photo.Caption, &photo.Path, &photo.Timestamp)
		if err != nil {
			if err == sql.ErrNoRows {
				// No rows found, return nil
				return nil, nil
			}
			// Some other error occurred
			return nil, fmt.Errorf("error fetching photo: %v", err)
		}

		return &photo, nil
	}
*/
func (db *appdbimpl) GetPhoto(userID int, photoID int) (*Photo, error) {
	query := `SELECT photo_id, user_id, photo, caption, timestamp
              FROM Photo
              WHERE user_id = ? AND photo_id = ?`

	// Log the SQL query and parameters
	fmt.Printf("Executing query: %s with userID=%d and photoID=%d\n", query, userID, photoID)

	// Execute the query
	row := db.c.QueryRow(query, userID, photoID)

	// Create a Photo instance to hold the result
	var photo Photo
	var imageData []byte

	// Scan the result into the Photo instance
	err := row.Scan(&photo.PhotoID, &photo.UserID, &imageData, &photo.Caption, &photo.Timestamp)
	if err != nil {
		if err == sql.ErrNoRows {
			// Log that no rows were found
			fmt.Printf("No photo found for userID=%d and photoID=%d\n", userID, photoID)
			// No rows found, return nil
			return nil, nil
		}
		// Log other errors
		fmt.Printf("Error fetching photo for userID=%d and photoID=%d: %v\n", userID, photoID, err)
		// Some other error occurred
		return nil, fmt.Errorf("error fetching photo: %v", err)
	}

	// Convert image data to base64-encoded string
	if imageData != nil {
		photo.Image = fmt.Sprintf("data:image/jpeg;base64,%s", base64.StdEncoding.EncodeToString(imageData))
	} else {
		photo.Image = "" // Or handle the case where image data is nil
	}

	return &photo, nil
}

func (db *appdbimpl) CheckPhoto(user_id int, photo_id int) (bool, error) {
	CHECK := `SELECT COUNT(*) FROM Photo WHERE user_id = ? AND photo_id = ?`

	var photoExists bool

	err := db.c.QueryRow(CHECK, user_id, photo_id).Scan(&photoExists)
	if err != nil || !photoExists {
		return false, err
	}

	return photoExists, err

}

func (db *appdbimpl) GetPhotoOwner(photoID int) (int, error) {
	query := `SELECT user_id FROM Photo WHERE photo_id = ?`
	var ownerID int

	err := db.c.QueryRow(query, photoID).Scan(&ownerID)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, fmt.Errorf("photo with id %d does not exist", photoID)
		}
		return 0, err
	}

	return ownerID, nil
}
