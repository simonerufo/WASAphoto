package api

import (
	"encoding/base64"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type User struct {
	UserID   int    `json:"id"`
	Username string `json:"username"`
}

type UserPhoto struct {
	PhotoID      int       `json:"photo_id"`
	UserID       int       `json:"user_id"`
	Image        string    `json:"image"`        // Base64-encoded image
	Timestamp    time.Time `json:"timestamp"`    // Date and time of upload
	LikeCount    int       `json:"likeCount"`    // Number of likes
	CommentCount int       `json:"commentCount"` // Number of comments
	Liked        bool      `json:"liked"`        // Whether the photo is liked by the user
	Caption      string    `json:"caption"`      // Caption for the photo
	Time         string    `json:"time"`         // Publication time
}

type Like struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	PhotoID  int    `json:"photo_id"`
}

type Comment struct {
	Username    string `json:"username"`
	CommentID   int    `json:"comment_id"`
	CommentText string `json:"comment_text"`
	Timestamp   string `json:"timestamp"`
}

type CommentRequest struct {
	CommentText string `json:"comment_text"`
}

func encodeToBase64(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

// checks if the username is valid
// func (user User) isNameValid() bool {
// 	rex := regexp.MustCompile(`^[a-z0-9]{3,13}$`) //regex to compile
// 	if !rex.MatchString(user.Username) {
// 		return false
// 	}
// 	return true
// }

// checks if the username is valid
func validUsername(username string) bool {
	rex := regexp.MustCompile(`^[a-z0-9]{3,13}$`) //regex to compile
	return rex.MatchString(username)
}

func Auth(w http.ResponseWriter, r *http.Request) (int, error) {
	authHeader := r.Header.Get("Authorization")
	// Expecting format "Bearer <token>"
	const bearerPrefix = "Bearer "
	if !strings.HasPrefix(authHeader, bearerPrefix) {
		http.Error(w, "Invalid Authorization header format", http.StatusBadRequest)
		return 0, http.ErrAbortHandler
	}

	// Extract token part
	tokenStr := strings.TrimPrefix(authHeader, bearerPrefix)
	userID, err := strconv.Atoi(tokenStr)
	if err != nil {
		http.Error(w, "Invalid Auth token", http.StatusBadRequest)
		return 0, err
	}

	return userID, nil
}
