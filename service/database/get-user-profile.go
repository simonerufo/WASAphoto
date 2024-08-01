package database

import( 
	"fmt"
	"encoding/base64"
	"database/sql"
)
/*
// Getting the User profile from his id
func (db *appdbimpl) GetUserProfile(userID int) (Profile, error) {
    var userProfile Profile
    var user User
    var followers int
    var following int

    // Fetch user profile using user_id
    user, err := db.GetUserByID(userID)
    if err != nil {
        return userProfile, err
    }

    // Fetch the number of followers
    getFollowersQuery := `SELECT COUNT(*) FROM Follow WHERE following_id = ?`
    err = db.c.QueryRow(getFollowersQuery, userID).Scan(&followers)
    if err != nil {
        return userProfile, err
    }

    // Fetch the number of following
    getFollowingQuery := `SELECT COUNT(*) FROM Follow WHERE followed_id = ?`
    err = db.c.QueryRow(getFollowingQuery, userID).Scan(&following)
    if err != nil {
        return userProfile, err
    }

    // Fetch photos
    var photos []Photo
    getPhotosQuery := `SELECT photo_id, user_id, photo, caption, timestamp, like_count, comment_count, liked, time
                       FROM Photo WHERE user_id = ?`
    rows, err := db.c.Query(getPhotosQuery, userID)
    if err != nil {
        return userProfile, err
    }
    defer rows.Close()

    for rows.Next() {
        var photo Photo
        var imageData []byte // To hold the raw image data
        
        err := rows.Scan(&photo.PhotoID, &photo.UserID, &imageData, &photo.Caption, &photo.Timestamp, &photo.LikeCount, &photo.CommentCount, &photo.Liked, &photo.Time)
        if err != nil {
            return userProfile, err
        }

        // Convert image data to base64-encoded string
        photo.Image = fmt.Sprintf("data:image/jpeg;base64,%s", base64.StdEncoding.EncodeToString(imageData))

        photos = append(photos, photo)
    }

    if err = rows.Err(); err != nil {
        return userProfile, err
    }

    // Populate profile struct
    userProfile.User = user
    userProfile.Photos = photos
    userProfile.Followers = followers
    userProfile.Following = following

    return userProfile, nil
}
*/
func (db *appdbimpl) GetUserProfileByUsername(username string) (Profile, error) {
    var userProfile Profile
    var user User
    var followers int
    var following int

    // Get user ID from username
    var userID int
    getUserIDQuery := `SELECT user_id FROM User WHERE username = ?`
    err := db.c.QueryRow(getUserIDQuery, username).Scan(&userID)
    if err != nil {
        if err == sql.ErrNoRows {
            return userProfile, fmt.Errorf("user not found")
        }
        return userProfile, err
    }

    // Get user profile details using the user ID
    user, err = db.GetUserByID(userID)
    if err != nil {
        return userProfile, err
    }

    // Get photos related to the user
    var photos []Photo
    getPhotosQuery := `SELECT photo_id, user_id, photo, caption, timestamp 
                       FROM Photo WHERE user_id = ?`
    rows, err := db.c.Query(getPhotosQuery, userID)
    if err != nil {
        return userProfile, err
    }
    defer rows.Close()

    for rows.Next() {
        var photo Photo
        var imageData []byte // To hold the raw image data
        
        err := rows.Scan(&photo.PhotoID, &photo.UserID, &imageData, &photo.Caption, &photo.Timestamp)
        if err != nil {
            return userProfile, err
        }

        // Convert image data to base64-encoded string
        photo.Image = fmt.Sprintf("data:image/jpeg;base64,%s", base64.StdEncoding.EncodeToString(imageData))

        photos = append(photos, photo)
    }

    if err = rows.Err(); err != nil {
        return userProfile, err
    }

    // Get follower and following counts
    getFollowersQuery := `SELECT COUNT(*) FROM Follow WHERE followed_id = ?`
    err = db.c.QueryRow(getFollowersQuery, userID).Scan(&followers)
    if err != nil {
        return userProfile, err
    }

    getFollowingQuery := `SELECT COUNT(*) FROM Follow WHERE following_id = ?`
    err = db.c.QueryRow(getFollowingQuery, userID).Scan(&following)
    if err != nil {
        return userProfile, err
    }

    // Populate the Profile struct
    userProfile.User = user
    userProfile.Photos = photos
    userProfile.Followers = followers
    userProfile.Following = following

    return userProfile, nil
}

func (db *appdbimpl) GetUserProfileByID(profileID int) (Profile, error) {
    var userProfile Profile
    var user User
    var followers, following int

    // Fetch user details
    getUserQuery := `SELECT user_id, username FROM User WHERE user_id = ?`
    err := db.c.QueryRow(getUserQuery, profileID).Scan(&user.UserID, &user.Username)
    if err != nil {
        fmt.Println("Error fetching user details:", err)
        return userProfile, err
    }

    // Fetch the number of followers
    getFollowersQuery := `SELECT COUNT(*) FROM Follow WHERE followed_id = ?`
    err = db.c.QueryRow(getFollowersQuery, profileID).Scan(&followers)
    if err != nil {
        fmt.Println("Error fetching followers count:", err)
        return userProfile, err
    }

    // Fetch the number of following
    getFollowingQuery := `SELECT COUNT(*) FROM Follow WHERE following_id = ?`
    err = db.c.QueryRow(getFollowingQuery, profileID).Scan(&following)
    if err != nil {
        fmt.Println("Error fetching following count:", err)
        return userProfile, err
    }

    // Fetch photos
    var photos []Photo
    getPhotosQuery := `SELECT photo_id, user_id, photo, caption, timestamp, 
                              (SELECT COUNT(*) FROM Like WHERE post_id = Photo.photo_id) AS likeCount,
                              (SELECT COUNT(*) FROM Comment WHERE post_id = Photo.photo_id) AS commentCount,
                              EXISTS(SELECT 1 FROM Like WHERE post_id = Photo.photo_id AND user_id = ?) AS liked
                       FROM Photo WHERE user_id = ?`
    rows, err := db.c.Query(getPhotosQuery, profileID, profileID)
    if err != nil {
        fmt.Println("Error fetching photos:", err)
        return userProfile, err
    }
    defer rows.Close()

    for rows.Next() {
        var photo Photo
        var imageData []byte

        err := rows.Scan(&photo.PhotoID, &photo.UserID, &imageData, &photo.Caption, &photo.Timestamp, &photo.LikeCount, &photo.CommentCount, &photo.Liked)
        if err != nil {
            fmt.Println("Error scanning photo row:", err)
            return userProfile, err
        }

        // Convert image data to base64-encoded string
        photo.Image = fmt.Sprintf("data:image/jpeg;base64,%s", base64.StdEncoding.EncodeToString(imageData))

        photos = append(photos, photo)
    }

    if err = rows.Err(); err != nil {
        fmt.Println("Error during rows iteration:", err)
        return userProfile, err
    }

    // Populate profile struct
    userProfile.User = user
    userProfile.Photos = photos
    userProfile.Followers = followers
    userProfile.Following = following

    return userProfile, nil
}




