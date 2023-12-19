package database

import (
	"fmt"
	"time"
)

func (db *appdbimpl) InsertPhoto(userID int, image []byte) (int, string, error) {

	date := time.Now().Format(time.RFC3339)

	result, err := db.c.Exec("INSERT INTO photos (userID, image, date) VALUES (?,?,?)", userID, image, date)
	if err != nil {
		fmt.Println("DATABASE: Error inserting image: ", err)
		return 0, date, err
	}

	photoID, err := result.LastInsertId()
	if err != nil {
		fmt.Println("DATABASE: Error obtaining the last insert id: ", err)
		return 0, date, err
	}

	return int(photoID), date, nil
}

func (db *appdbimpl) GetFollowingByUserID(userID int) ([]int, error) {

	var following []int

	rows, err := db.c.Query("SELECT followID FROM follow WHERE id = ?", userID)
	if err != nil {
		fmt.Println("DATABASE: Error obtaining user ids of the users followed by the user: ", err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var followID int
		if err := rows.Scan(&followID); err != nil {
			return nil, err
		}
		following = append(following, followID)
	}
	return following, nil
}

// Manca reverse e foto completa
func (db *appdbimpl) GetStream(userID int) ([]Photo, error) {
	var stream []Photo

	following, err := db.GetFollowingByUserID(userID)
	if err != nil {
		return nil, err
	}

	// Query all photos from users followed by the given userID
	query := fmt.Sprintf("SELECT * FROM photos WHERE userID IN (%s) ORDER BY date DESC", joinIntSlice(following))
	rows, err := db.c.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var photo Photo
		if err := rows.Scan(&photo.PhotoID, &photo.UserID, &photo.Image, &photo.Date); err != nil {
			return nil, err
		}
		photo.Image = nil
		stream = append(stream, photo)
	}

	return stream, nil
}

// Utility function to join an int slice into a comma-separated string
func joinIntSlice(slice []int) string {
	result := ""
	for i, id := range slice {
		result += fmt.Sprintf("%d", id)
		if i < len(slice)-1 {
			result += ","
		}
	}
	return result
}
