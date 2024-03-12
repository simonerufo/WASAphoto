package database

import "time"

type User struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
}

type Post struct {
	PostID      int       `json:"post_id"`
	User        User      `json:"user"`
	Caption     string    `json:"caption"`
	ImageURL    string    `json:"imageURL"`
	LikesNum    int       `json:"likesNum"`
	CommentsNum int       `json:"commentsNum"`
	Timestamp   time.Time `json:"timestamp"`
}

type Comment struct {
	CommentID int       `json:"comment_id"`
	PostID    int       `json:"post_id"`
	OwnerID   int       `json:"owner_id"`
	User      User      `json:"user_id"`
	Text      string    `json:"text"`
	Timestamp time.Time `json:"timestamp"`
}

type Photo struct {
	UserID    int       `json:"user_id"`
	PhotoID   int       `json:"pid"`
	Path      string    `json:"photo_path"`
	Caption   string    `json:"caption"`
	Timestamp time.Time `json:"time"`
}
type Profile struct {
	User User `json:"user"`
	//Photos    []Photo `json:"photos"`
	Followers int `json:"followers"`
	Following int `json:"following"`
}
