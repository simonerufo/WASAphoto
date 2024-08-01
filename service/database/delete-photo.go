package database
import ("fmt")

func (db *appdbimpl) DeletePhoto(user_id int, photo_id int) error {
	DELETE := `DELETE FROM Photo WHERE user_id = ? AND photo_id = ?`

	_, err := db.c.Exec(DELETE, user_id, photo_id)
	if err != nil {
		fmt.Printf("Error while trying to delete photo")
		return err
	}

	return err
}
