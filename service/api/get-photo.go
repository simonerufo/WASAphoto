package api
import( 
	"net/http"
	"encoding/json"
	"fmt" 
    "time"
	"strconv"
	"github.com/julienschmidt/httprouter"
    "git.simonerufo.it/WASAphoto/service/api/reqcontext"
)
/*
func (rt *_router) getPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
    // Extract the user ID and photo ID from the URL parameters
    userIDStr := ps.ByName("user_id")
    photoIDStr := ps.ByName("photo_id")
    
    // Convert userID and photoID to integers
    userID, err := strconv.Atoi(userIDStr)
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }
    
    photoID, err := strconv.Atoi(photoIDStr)
    if err != nil {
        http.Error(w, "Invalid photo ID", http.StatusBadRequest)
        return
    }
    
    // Fetch the photo details from the database
    photo, err := rt.db.GetPhoto(userID, photoID)
    if err != nil {
        fmt.Printf("Error fetching photo for user ID %d and photo ID %d: %v\n", userID, photoID, err)
        http.Error(w, "Error fetching photo", http.StatusInternalServerError)
        return
    }

    // Check if the photo is found
    if photo == nil {
        http.Error(w, "Photo not found", http.StatusNotFound)
        return
    }

    // Set response headers and encode the response
    w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(photo); err != nil {
        http.Error(w, "Error encoding response", http.StatusInternalServerError)
        return
    }

    fmt.Printf("PHOTO: user_id=%d, photo_id=%d\n", userID, photoID)
}
*/
func (rt *_router) getPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
    // Extract the user ID and photo ID from the URL parameters
    userIDStr := ps.ByName("user_id")
    photoIDStr := ps.ByName("photo_id")

    // Log the received parameters
    fmt.Printf("Received GET request with user_id=%s and photo_id=%s\n", userIDStr, photoIDStr)
    
    // Convert userID and photoID to integers
    userID, err := strconv.Atoi(userIDStr)
    if err != nil {
        fmt.Printf("Error converting user_id to integer: %v\n", err)
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }

    photoID, err := strconv.Atoi(photoIDStr)
    if err != nil {
        fmt.Printf("Error converting photo_id to integer: %v\n", err)
        http.Error(w, "Invalid photo ID", http.StatusBadRequest)
        return
    }

    // Log the converted IDs
    fmt.Printf("Converted user_id=%d and photo_id=%d\n", userID, photoID)
    
    // Fetch the photo details from the database
    photoData, err := rt.db.GetPhoto(userID, photoID)
    if err != nil {
        fmt.Printf("Error fetching photo for user_id=%d and photo_id=%d: %v\n", userID, photoID, err)
        http.Error(w, "Error fetching photo", http.StatusInternalServerError)
        return
    }

    // Check if the photo is found
    if photoData == nil {
        fmt.Printf("Photo not found for user_id=%d and photo_id=%d\n", userID, photoID)
        http.Error(w, "Photo not found", http.StatusNotFound)
        return
    }

    // Convert photo data to UserPhoto struct
    photo := UserPhoto{
        PhotoID:     photoData.PhotoID,
        UserID:      photoData.UserID,
        Image:       photoData.Image,
        Timestamp:   photoData.Timestamp,
        Caption:     photoData.Caption,
        Time:        photoData.Timestamp.Format(time.RFC3339),
        LikeCount:   0,
        CommentCount: 0,
        Liked:       false, 
    }

    // Log the retrieved photo details
    fmt.Printf("Retrieved photo: %+v\n", photo)

    // Set response headers and encode the response
    w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(photo); err != nil {
        fmt.Printf("Error encoding response: %v\n", err)
        http.Error(w, "Error encoding response", http.StatusInternalServerError)
        return
    }

    // Log successful response
    fmt.Printf("Successfully returned photo for user_id=%d and photo_id=%d\n", userID, photoID)
}



