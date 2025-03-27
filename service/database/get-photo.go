package database

import (
	"database/sql"
	"encoding/base64"
	"errors"
	"log"
)

func (db *appdbimpl) GetPhoto(userID int, photoID int) (*Photo, error) {
	query := `SELECT photo_id, user_id, photo, caption, timestamp
              FROM Photo
              WHERE user_id = ? AND photo_id = ?`

	// Execute the query
	row := db.c.QueryRow(query, userID, photoID)

	// Create a Photo instance to hold the result
	var photo Photo
	var imageData []byte

	// Scan the result into the Photo instance
	err := row.Scan(&photo.PhotoID, &photo.UserID, &imageData, &photo.Caption, &photo.Timestamp)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// Log that no rows were found
			log.Printf("No photo found for userID=%d and photoID=%d", userID, photoID)
			// No rows found, return nil
			return nil, nil
		}
		// Log other errors
		log.Printf("Error fetching photo for userID=%d and photoID=%d: %v", userID, photoID, err)
		// Some other error occurred
		return nil, errors.New("failed to fetch photo")
	}

	// Convert image data to base64-encoded string
	if imageData != nil {
		photo.Image = "data:image/jpeg;base64," + base64.StdEncoding.EncodeToString(imageData)
	} else {
		photo.Image = ""
	}

	return &photo, nil
}

// checkPhoto checks if a certain photo_id associated to an user_id exists
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
		if errors.Is(err, sql.ErrNoRows) {
			log.Printf("Photo with id %d does not exist", photoID)
			return 0, errors.New("photo does not exist")
		}
		log.Printf("Error retrieving owner for photoID %d: %v", photoID, err)
		return 0, errors.New("failed to retrieve photo owner")
	}

	return ownerID, nil
}
