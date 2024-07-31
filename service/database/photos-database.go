package database

import (
	"database/sql"
	"errors"
	"time"
)

func (db *appdbimpl) PhotoExists(photoID int, photoAuthorID int) bool {

	var id int
	err := db.c.QueryRow("SELECT id FROM photos WHERE id = ? AND userID = ?", photoID, photoAuthorID).Scan(&id)
	return err == nil
}

func (db *appdbimpl) InsertPhoto(userID int, image []byte) (int, string, error) {

	date := time.Now().Format(time.RFC3339)

	result, err := db.c.Exec("INSERT INTO photos (userID, image, date) VALUES (?,?,?)", userID, image, date)
	if err != nil {
		return 0, date, err
	}

	photoID, err := result.LastInsertId()
	if err != nil {
		return 0, date, err
	}

	return int(photoID), date, nil
}

func (db *appdbimpl) DeletePhoto(photoID int) error {

	_, err := db.c.Exec("DELETE FROM comments WHERE photoID = ?", photoID)
	if err != nil {
		return err
	}

	_, err = db.c.Exec("DELETE FROM likes WHERE photoID = ?", photoID)
	if err != nil {
		return err
	}

	result, err := db.c.Exec("DELETE FROM photos WHERE id = ?", photoID)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("photo not found")
	}

	return nil
}

func (db *appdbimpl) GetCompletePhotos(rows *sql.Rows) ([]Photo, error) {

	var stream []Photo

	for rows.Next() {
		var photo Photo
		if err := rows.Scan(&photo.PhotoID, &photo.UserID, &photo.Image, &photo.Date, &photo.Username); err != nil {
			return nil, err
		}
		photo.Comments, _ = db.GetComments(photo.PhotoID)
		photo.Likes, _ = db.GetLikes(photo.PhotoID)
		stream = append(stream, photo)
	}

	// Check errors during iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}

	if len(stream) == 0 {
		return []Photo{}, nil
	}

	return stream, nil
}

func (db *appdbimpl) GetUserPhotos(userID int) ([]Photo, error) {

	query := `SELECT photos.*, users.username
			  FROM photos 
			  JOIN users ON photos.userID = users.id
			  WHERE photos.userID = ? 
			  ORDER BY date DESC`

	rows, err := db.c.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return db.GetCompletePhotos(rows)
}

func (db *appdbimpl) GetStream(userID int) ([]Photo, error) {

	// Query all photos from users followed by the given userID
	query := `SELECT photos.*, users.username
			  FROM photos
			  JOIN users ON photos.userID = users.id
			  WHERE photos.userID IN (SELECT followID FROM follow WHERE userID = ?)
			  ORDER BY date DESC`

	rows, err := db.c.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return db.GetCompletePhotos(rows)
}
