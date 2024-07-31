package database

import("fmt")
// InsertPhoto inserts a new photo into the Photo table with a manually managed photo_id and returns the photo ID.
func (db *appdbimpl) InsertPhoto(userID int, caption string, photo []byte) (int, error) {
    // Fetch the current maximum photo_id
    var maxID int
    err := db.c.QueryRow(`SELECT COALESCE(MAX(photo_id), 0) FROM Photo`).Scan(&maxID)
    if err != nil {
        return 0, fmt.Errorf("error fetching max photo_id: %w", err)
    }

    // Increment the ID
    newPhotoID := maxID + 1

    //Insert the new photo record with the incremented photo_id
    UPLOAD := `INSERT INTO Photo(photo_id, user_id, caption, photo) 
               VALUES (?, ?, ?, ?)`

    _, err = db.c.Exec(UPLOAD, newPhotoID, userID, caption, photo)
    if err != nil {
        return 0, fmt.Errorf("error inserting photo: %w", err)
    }

    // Log the insertion
    fmt.Printf("Inserted photo with ID %d for user %d\n", newPhotoID, userID)

    // Return the new photo ID
    return newPhotoID, nil
}
