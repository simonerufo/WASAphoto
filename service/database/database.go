/*
Package database is the middleware between the app database and the code. All data (de)serialization (save/load) from a
persistent database are handled here. Database specific logic should never escape this package.

To use this package you need to apply migrations to the database if needed/wanted, connect to it (using the database
data source name from config), and then initialize an instance of AppDatabase from the DB connection.

For example, this code adds a parameter in `webapi` executable for the database data source name (add it to the
main.WebAPIConfiguration structure):

	DB struct {
		Filename string `conf:""`
	}

This is an example on how to migrate the DB and connect to it:

	// Start Database
	logger.Println("initializing database support")
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		logger.WithError(err).Error("error opening SQLite DB")
		return fmt.Errorf("opening SQLite: %w", err)
	}
	defer func() {
		logger.Debug("database stopping")
		_ = db.Close()
	}()

Then you can initialize the AppDatabase and pass it to the api package.
*/
package database

import (
	"database/sql"
	"errors"
	"fmt"
)

// AppDatabase is the high level interface for the DB
type AppDatabase interface {
	//GetName() (string, error)
	//SetName(name string) error

	//insert user in user table
	InsertUser(username string) error
	//search user from user table
	SearchUsername(username string) bool
	//get user from user table
	GetUserByName(username string) (User, error)
	GetUserByID(id int) (User, error)
	//update username in user table
	UpdateName(user User) error
	//add user follow relation
	FollowUser(following_id int, followed_id int) error
	//remove a user follow relation
	UnfollowUser(following_id int, followed_id int) error
	//checks if a user banned another user
	GetBan(banning_id int, banned_id int) (bool, error)
	//remove a user ban relation
	UnbanUser(banning_id int, banned_id int) error
	//checks if a user followed another user
	GetFollow(following_id int, followed_id int) (bool, error)
	//retrieves all users that followed the user with userID
	GetFollowersForUser(userID int) ([]User, error)
	//retrieves all followed user by using his userID
	GetFollowingForUser(userID int) ([]User, error)
	//ban a user adding him in db
	BanUser(banning_id int, banned_id int) error
	//getting from db user profile using id
	GetUserProfileByID(profileID int) (Profile, error)
	//getting from db user profile using username
	GetUserProfileByUsername(username string) (Profile, error)
	//retrieving the user stream
	GetStream(user_id int) ([]Photo, error)
	//uploading a photo(post) into db
	InsertPhoto(userID int, caption string, photo []byte) (int, error)
	//checks if a photo exists in db
	CheckPhoto(user_id int, photo_id int) (bool, error)
	//gets a photo from db
	GetPhoto(userID int, photoID int) (*Photo, error)
	//gets the owner of a photo by using his corresponding photo id
	GetPhotoOwner(photoID int) (int, error)
	//deletes a photo from db
	DeletePhoto(user_id int, photo_id int) error
	//insert an entry into like table
	InsertLike(user_id int, owner_id int, post_id int) error
	//delete an entry from like table
	DeleteLike(user_id int, owner_id int, post_id int) error
	//Get all likes related to a photo
	GetLikesForPhoto(photoID int) ([]Like, error)
	//insert an entry into comment table
	InsertComment(user_id int, owner_id int, photo_id int, text string) (int, error)
	//check if a comment exists in comment table
	CheckComment(user_id int, photo_id int, comment_id int) (bool, error)
	//Getting all comments referred to a photo
	GetCommentsByPhotoID(postID int) ([]Comment, error)
	//Get a comment using his id
	GetCommentByID(commentID int) (Comment, error)
	//delete a comment from comments table
	DeleteComment(user_id int, photo_id int, comment_id int) error
	//retrieves the photo id from a comment and user id
	GetPhotoIDFromComment(commentID int, userID int) (int, error)
	// retrieves the photo id from the comment and user id of the post owner
	GetPhotoIDByOwner(commentID int, userID int) (int, error)
	// checks if a specific user owns the photo referred to a specific comment
	CheckPhotoOwnership(commentID int, userID int) (bool, error)
	// retrieves the ban list of a user
	GetBanList(userID int) ([]int, error)

	Ping() error
}

type appdbimpl struct {
	c *sql.DB
}

// New returns a new instance of AppDatabase based on the SQLite connection `db`.
// `db` is required - an error will be returned if `db` is `nil`.
func New(db *sql.DB) (AppDatabase, error) {
	if db == nil {
		return nil, errors.New("database is required when building a AppDatabase")
	}

	// Check if table exists. If not, the database is empty, and we need to create the structure
	var tablesNumber int
	err := db.QueryRow(`SELECT count(name) FROM sqlite_master WHERE type='table';`).Scan(&tablesNumber)
	/*if errors.Is(err, sql.ErrNoRows) {
	sqlStmt := `CREATE TABLE example_table (id INTEGER NOT NULL PRIMARY KEY, name TEXT);`
	_, err = db.Exec(sqlStmt)*/
	if err != nil {
		return nil, fmt.Errorf("error creating database structure: %w", err)
	}
	//}
	if tablesNumber != 6 {
		//USER_TABLE
		_, err = db.Exec(USER_TABLE)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure (USER): %w", err)
		}

		//POST_TABLE
		_, err = db.Exec(PHOTO_TABLE)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure (PHOTO): %w", err)
		}

		//LIKE_TABLE
		_, err = db.Exec(LIKE_TABLE)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure (LIKE): %w", err)
		}

		//COMMENT_TABLE
		_, err = db.Exec(COMMENT_TABLE)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure (COMMENT): %w", err)
		}

		//FOLLOW_TABLE
		_, err = db.Exec(FOLLOW_TABLE)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure (FOLLOW): %w", err)
		}

		//BAN_TABLE
		_, err = db.Exec(BAN_TABLE)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure (BAN): %w", err)
		}
	}

	return &appdbimpl{
		c: db,
	}, nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
