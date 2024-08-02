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

// type Comment struct {
// 	CommentID int       `json:"comment_id"`
// 	PostID    int       `json:"post_id"`
// 	OwnerID   int       `json:"owner_id"`
// 	User      User      `json:"user_id"`
// 	Text      string    `json:"text"`
// 	Timestamp time.Time `json:"timestamp"`
// }

type Comment struct {
	Username    string `json:"username"`
	CommentID   int    `json:"comment_id"`
	CommentText string `json:"comment_text"`
	Timestamp   string `json:"timestamp"`
}

type Photo struct {
	PhotoID		int   `json:"photo_id"`
	UserID      int  `json:"user_id"`
	Image       string    `json:"image"`       // Base64-encoded image
	Timestamp   time.Time `json:"timestamp"`   // Date and time of upload
	LikeCount   int       `json:"likeCount"`   // Number of likes
	CommentCount int      `json:"commentCount"` // Number of comments
	Liked       bool      `json:"liked"`       // Whether the photo is liked by the user
	Caption     string    `json:"caption"`     // Caption for the photo
	Time        string    `json:"time"`        // Publication time
}

type Profile struct {
	User User `json:"user"`
	Photos    []Photo `json:"photos"`
	Followers int `json:"followers"`
	Following int `json:"following"`
}

