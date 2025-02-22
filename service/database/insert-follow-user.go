package database

import (
	"errors"
	"log"
)

// insert into Follow tables users ids of user that wants to follow another
// func (db *appdbimpl) FollowUser(following_id int, followed_id int) error {
// 	FOLLOW := `INSERT INTO Follow(following_id,followed_id)
// 				VALUES(?,?)`

// 	_, err := db.c.Exec(FOLLOW, following_id, followed_id)
// 	if err != nil {
// 		return err
// 	}

//		return err
//	}
func (db *appdbimpl) FollowUser(following_id int, followed_id int) error {
	// Check if the follow relationship already exists
	var exists bool
	CHECK := `SELECT EXISTS(SELECT 1 FROM Follow WHERE following_id = ? AND followed_id = ?)`
	err := db.c.QueryRow(CHECK, following_id, followed_id).Scan(&exists)
	if err != nil {
		log.Println("DB Check Error:", err)
		return err
	}

	if exists {
		return errors.New("user already follows this person")
	}

	// Insert the new follow relationship
	FOLLOW := `INSERT INTO Follow(following_id, followed_id) VALUES(?, ?)`
	_, err = db.c.Exec(FOLLOW, following_id, followed_id)
	if err != nil {
		log.Println("DB Insert Error:", err)
	}
	return err
}
