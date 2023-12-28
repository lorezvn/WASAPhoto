package database

import (
	"errors"
	"time"
)

func (db *appdbimpl) InsertLike(userID int, photoID int) (string, error) {

	date := time.Now().Format(time.RFC3339)

	_, err := db.c.Exec("INSERT INTO likes (userID, photoID, date) VALUES (?,?,?)", userID, photoID, date)
	if err != nil {
		return date, err
	}

	return date, nil
}

func (db *appdbimpl) DeleteLike(userID int, photoID int) error {
	result, err := db.c.Exec("DELETE FROM likes WHERE userID = ? AND photoID = ?", userID, photoID)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("Like not found")
	}

	return nil
}

func (db *appdbimpl) GetLikes(photoID int) ([]Like, error) {

	var likes []Like
	rows, err := db.c.Query("SELECT * FROM likes WHERE photoID = ?", photoID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var like Like
		if err := rows.Scan(&like.UserID, &like.PhotoID, &like.Date); err != nil {
			return nil, err
		}

		likes = append(likes, like)
	}

	return likes, nil
}