package api

import (
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"fmt" 	
	"time"
	"encoding/base64"
)

type User struct {
	UserID   int    `json:"id"`
	Username string `json:"username"`
}

type UserPhoto struct {
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

func encodeToBase64(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}
// checks if the username is valid
func (user User) isNameValid() bool {
	rex := regexp.MustCompile(`^[a-z0-9]{3,13}$`) //regex to compile
	if !rex.MatchString(user.Username) {
		return false
	}
	return true
}

// checks if the username is valid
func validUsername(username string) bool {
	rex := regexp.MustCompile(`^[a-z0-9]{3,13}$`) //regex to compile
	return rex.MatchString(username)
}

/*
// check if user is authorized
func isAuth(w http.ResponseWriter, r *http.Request, token int) bool {
	//getting the Authorization Header from the http request and cast it to int
	auth, err := strconv.Atoi(r.Header.Get("Authorization"))
	if err != nil {
		http.Error(w, "cannot convert properly Authorization header", http.StatusBadRequest)
		return false
	}

	if token != auth {
		http.Error(w, "not authorized", http.StatusBadRequest)
		return false
	}
	return true
}
*/
func isAuth(w http.ResponseWriter, r *http.Request, id ...int) (int, error) {
	auth := strings.Split(r.Header.Get("Authorization"), " ")
	
	if len(auth) <= 1 {
		return 0, nil
	}
	uid, err := strconv.Atoi(auth[1])
	fmt.Printf("auth:%d\n",uid)
	if err != nil || (len(id) > 0 && uid != id[0]) {
		http.Error(w, "Unauthorized token", http.StatusUnauthorized)
		return 0, err
	}
	return uid, nil
}



func Auth(w http.ResponseWriter, r *http.Request) (int, error) {
	// Retrieving the Authorization header
	authHeader := r.Header.Get("Authorization")
	// Expecting format "Bearer <token>"
	const bearerPrefix = "Bearer "
	if !strings.HasPrefix(authHeader, bearerPrefix) {
		fmt.Printf("Invalid Authorization header format. Expected prefix: '%s'\n", bearerPrefix)
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
