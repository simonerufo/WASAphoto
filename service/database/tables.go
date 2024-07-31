package database

// USER
/*
	-id: identificator for user
	-username: name assigned to the user
*/

var USER_TABLE = `CREATE TABLE IF NOT EXISTS User(
	user_id INTEGER PRIMARY KEY AUTOINCREMENT,
	username STRING NOT NULL UNIQUE,
	UNIQUE (user_id) ON CONFLICT REPLACE
);`

/*
// adding some entries in user table
var INSERT_USERS = `INSERT INTO User(username)
					VALUES
					("simone"),
					("gigi"),
					("diana"),
					("daniele"),
					("andrea");`
*/

// PHOTO
/*



var PHOTO_TABLE = `CREATE TABLE IF NOT EXISTS Photo(
	photo_id INTEGER NOT NULL,
	user_id INTEGER NOT NULL,
	caption STRING NOT NULL,
	photo BLOB NOT NULL,
	timestamp DATETIME DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY(photo_id,user_id),
	FOREIGN KEY (user_id) REFERENCES User(user_id)
		ON DELETE CASCADE
);`
*/
var PHOTO_TABLE = `CREATE TABLE IF NOT EXISTS Photo(
    photo_id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INT NOT NULL,
    caption STRING NOT NULL,
    photo BLOB NOT NULL,
    timestamp DATETIME DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(photo_id, user_id),
    FOREIGN KEY (user_id) REFERENCES User(user_id) ON DELETE CASCADE
);`

//COMMENT
/*
 	-comment_id: identificator of comment
	-post_id: identificator of post
	-owner_id: identificato of owner of the post
	-user_id: identificator of user that wants to post a comment
	-text : text wrote by user in the comment
	-timestamp: current timestamp assigned to comment whenever it's created
*/
var COMMENT_TABLE = `CREATE TABLE IF NOT EXISTS Comment(
	comment_id INTEGER AUTO_INCREMENT,
	post_id INTEGER NOT NULL,
	owner_id INTEGER NOT NULL,
	user_id INTEGER NOT NULL,
	text TEXT NOT NULL,
	timestamp DATETIME DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY(comment_id,post_id,owner_id)
	FOREIGN KEY(post_id,owner_id) REFERENCES Post(user_id,post_id)
		ON DELETE CASCADE,
	FOREIGN KEY(user_id) REFERENCES User(user_id)
		ON DELETE CASCADE
);`

//LIKE
/*
	-user_id : identificator of user that wants to like a post
	-owner_id: identificator of user that's receiving the like
	-post_id: identificator of post that's receiving the like
*/

var LIKE_TABLE = `CREATE TABLE IF NOT EXISTS Like(
	user_id INTEGER NOT NULL,
	owner_id INTEGER NOT NULL,
	post_id INTEGER NOT NULL,
	PRIMARY KEY(post_id,owner_id,user_id),
	FOREIGN KEY(post_id,owner_id) REFERENCES Post(post_id,user_id)
		ON DELETE CASCADE,
	FOREIGN KEY(user_id) REFERENCES Post(user_id)
		ON DELETE CASCADE
);`

//FOLLOW
/*
	-following_id: identificator of user that wants to follow another user
	-followed_id: identificator of user that's going to be followed by another user
*/

var FOLLOW_TABLE = `CREATE TABLE IF NOT EXISTS Follow(
	following_id INTEGER NOT NULL,
	followed_id INTEGER NOT NULL,
	PRIMARY KEY(following_id,followed_id),
	FOREIGN KEY(following_id) REFERENCES User(user_id)
		ON DELETE CASCADE,
	FOREIGN KEY(followed_id) REFERENCES User(user_id)
		ON DELETE CASCADE
);`

//BAN
/*
	-banning_id: identificator of user that wants to ban another user
	-banned_id: identificator of user that's going to be banned by another user
*/
var BAN_TABLE = `CREATE TABLE IF NOT EXISTS Ban(
	banning_id INTEGER NOT NULL,
	banned_id INTEGER NOT NULL,
	PRIMARY KEY(banning_id,banned_id)
	FOREIGN KEY(banning_id) REFERENCES User(user_id)
		ON DELETE CASCADE,
	FOREIGN KEY(banned_id) REFERENCES User(user_id)
		ON DELETE CASCADE
);`
