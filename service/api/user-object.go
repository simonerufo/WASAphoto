package api

import (
	"net/http"
	"regexp"
	"strconv"
)

type User struct {
	UserID   int    `json:"id"`
	Username string `json:"username"`
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
