package api

import (
	"fmt"
	"encoding/json"
	"io"
	"net/http"
	"time"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Retrieve user ID from parameters
	userID, err := strconv.Atoi(ps.ByName("user_id"))
	if err != nil {
		http.Error(w, "Error while retrieving user ID from parameters", http.StatusBadRequest)
		return
	}

	// Auth check
	_,err = Auth(w,r)
	if err != nil{
		http.Error(w,"Invalid authorization",http.StatusUnauthorized)
		return
	}

	// Parse the multipart form with a maximum file size limit
	err = r.ParseMultipartForm(16 << 20) // max 16 MB
	if err != nil {
		http.Error(w, "Error while parsing multipart form", http.StatusBadRequest)
		return
	}

	// Get the caption
	caption := r.FormValue("caption")

	// Get the image file
	image, _, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Error while parsing the image", http.StatusBadRequest)
		return
	}
	defer image.Close()

	// Read the image data
	imageData, err := io.ReadAll(image)
	if err != nil {
		http.Error(w, "Error while reading the image", http.StatusBadRequest)
		return
	}

	// Detect and validate the file type
	fileType := http.DetectContentType(imageData)
	if fileType != "image/jpeg" && fileType != "image/png" {
		http.Error(w, "Wrong file type", http.StatusBadRequest)
		return
	}

	// Insert the photo into the database and get the photo ID
	photoID, err := rt.db.InsertPhoto(userID, caption, imageData)
	if err != nil {
		http.Error(w, "Error while trying to insert photo into database", http.StatusInternalServerError)
		return
	}


	// Create UserPhoto response
	response := UserPhoto{
		PhotoID:     photoID,
		UserID:      userID,
		Image:       fmt.Sprintf("data:image/jpeg;base64,%s", encodeToBase64(imageData)), // Assuming image is JPEG
		Timestamp:   time.Now(),
		LikeCount:   0, 
		CommentCount: 0,
		Liked:       false, 
		Caption:     caption,
		Time:        time.Now().Format(time.RFC3339),
	}

	// Set response headers and write JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "error while encoding the json", http.StatusInternalServerError)
	}
}