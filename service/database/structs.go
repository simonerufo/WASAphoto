package database

import "time"

type User struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
}

type Post struct {
	PostID      int       `json:"post_id"`
	User        User      `json:"user"`
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
