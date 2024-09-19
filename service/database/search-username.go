package database

import (
	"database/sql"
	"log"
)

func (db *appdbimpl) SearchUsername(username string) bool {
	if username == "" {
		log.Println("Invalid input: username cannot be empty")
		return false
	}

	SEARCH := `SELECT username FROM User WHERE username = ?`
	var validUser string

	err := db.c.QueryRow(SEARCH, username).Scan(&validUser)

	if err != nil {
		if err == sql.ErrNoRows {
			// No user found with the given username
			return false
		}
		// Log unexpected errors for further investigation
		log.Printf("Error querying database for username %s: %v\n", username, err)
		return false
	}

	return true
}
