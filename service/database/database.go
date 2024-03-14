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
	//checks if a user banned another user
	GetBan(banning_id int, banned_id int) (bool, error)
	//getting from db user profile
	GetUserProfile(user_id int) (Profile, error)
	//retrieving the user stream
	GetStream(user_id int) ([]Post, error)

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
