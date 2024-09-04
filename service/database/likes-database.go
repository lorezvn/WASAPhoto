package database

import (
	"time"
)

func (db *appdbimpl) LikeExists(userID int, photoID int) bool {

	var count int
	err := db.c.QueryRow("SELECT COUNT(*) FROM likes WHERE userID = ? AND photoID = ?", userID, photoID).Scan(&count)
	if err != nil {
		return false
	}

	return count > 0
}

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
		return ErrLikeNotFound
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

	// Check errors during iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}

	if len(likes) == 0 {
		return []Like{}, nil
	}

	return likes, nil
}
