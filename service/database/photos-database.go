package database

import (
	"database/sql"
	"fmt"
	"time"
)

func (db *appdbimpl) PhotoExists(photoID int, photoAuthorID int) bool {

	var id int
	err := db.c.QueryRow("SELECT id FROM photos WHERE id = ? AND userID = ?", photoID, photoAuthorID).Scan(&id)
	if err != nil {
		return false
	}
	return true
}

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

func (db *appdbimpl) DeletePhoto(photoID int) error {
	_, err := db.c.Exec("DELETE FROM photos WHERE id = ?", photoID)
	_, err = db.c.Exec("DELETE FROM comments WHERE photoID = ?", photoID)
	_, err = db.c.Exec("DELETE FROM likes WHERE photoID = ?", photoID)
	return err
}

func (db *appdbimpl) InsertComment(userID int, photoID int, message string) (int, string, error) {

	date := time.Now().Format(time.RFC3339)

	result, err := db.c.Exec("INSERT INTO comments (userID, photoID, message, date) VALUES (?,?,?,?)", userID, photoID, message, date)
	if err != nil {
		fmt.Println("DATABASE: Error inserting comment: ", err)
		return 0, date, err
	}

	commentID, err := result.LastInsertId()
	if err != nil {
		fmt.Println("DATABASE: Error obtaining the last insert id: ", err)
		return 0, date, err
	}

	return int(commentID), date, nil
}

func (db *appdbimpl) DeleteComment(commentID int) (error) {
	_, err := db.c.Exec("DELETE FROM comments WHERE id = ?", commentID)
	return err
}

func (db *appdbimpl) InsertLike(userID int, photoID int) (string, error) {

	date := time.Now().Format(time.RFC3339)

	_, err := db.c.Exec("INSERT INTO likes (userID, photoID, date) VALUES (?,?,?)", userID, photoID, date)
	if err != nil {
		fmt.Println("DATABASE: Error inserting like: ", err)
		return date, err
	}

	return date, nil
}

func (db *appdbimpl) DeleteLike(likeID int) (error) {
	_, err := db.c.Exec("DELETE FROM likes WHERE id = ?", likeID)
	return err
}

func (db *appdbimpl) GetComments(photoID int) ([]Comment, error) {

	var comments []Comment
	rows, err := db.c.Query("SELECT * FROM comments WHERE photoID = ?", photoID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var comment Comment
		if err := rows.Scan(&comment.CommentID, &comment.UserID, &comment.PhotoID, &comment.Message, &comment.Date); err != nil {
			return nil, err
		}

		comments = append(comments, comment)
	}

	return comments, nil
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

func (db *appdbimpl) GetCompletePhotos(rows *sql.Rows) ([]Photo, error) {

	var stream []Photo

	for rows.Next() {
		var photo Photo
		if err := rows.Scan(&photo.PhotoID, &photo.UserID, &photo.Image, &photo.Date); err != nil {
			fmt.Println(err)
			return nil, err
		}
		photo.Image = nil
		comments, _ := db.GetComments(photo.PhotoID)
		likes, _ := db.GetLikes(photo.PhotoID)
		photo.Comments = comments
		photo.Likes = likes
		stream = append(stream, photo)
	}
	return stream, nil
}

func (db *appdbimpl) GetUserPhotos(userID int) ([]Photo, error) {

	// Query all photos from users followed by the given userID
	query := `SELECT * 
			  FROM photos 
			  WHERE photos.userID = ? 
			  ORDER BY date DESC`

	rows, err := db.c.Query(query, userID)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	return db.GetCompletePhotos(rows)
}

func (db *appdbimpl) GetStream(userID int) ([]Photo, error) {

	// Query all photos from users followed by the given userID
	query := `SELECT *
			  FROM photos
			  WHERE photos.userID IN (SELECT followID FROM follow WHERE id = ?)`

	rows, err := db.c.Query(query, userID)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	return db.GetCompletePhotos(rows)
}


