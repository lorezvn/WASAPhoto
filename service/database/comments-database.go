package database

import (
	"time"
)

func (db *appdbimpl) InsertComment(userID int, photoID int, message string) (int, string, error) {

	date := time.Now().Format(time.RFC3339)

	result, err := db.c.Exec("INSERT INTO comments (userID, photoID, message, date) VALUES (?,?,?,?)", userID, photoID, message, date)
	if err != nil {
		return 0, date, err
	}

	commentID, err := result.LastInsertId()
	if err != nil {
		return 0, date, err
	}

	return int(commentID), date, nil
}

func (db *appdbimpl) DeleteComment(commentID int) error {
	result, err := db.c.Exec("DELETE FROM comments WHERE id = ?", commentID)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return ErrCommentNotFound
	}

	return nil
}

func (db *appdbimpl) GetComments(photoID int) ([]Comment, error) {

	var comments []Comment

	query := `SELECT comments.*, users.username
			  FROM comments 
			  JOIN users ON comments.userID = users.id
			  WHERE photoID = ?`

	rows, err := db.c.Query(query, photoID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var comment Comment
		if err := rows.Scan(&comment.CommentID, &comment.UserID, &comment.PhotoID, &comment.Message, &comment.Date, &comment.Username); err != nil {
			return nil, err
		}

		comments = append(comments, comment)
	}

	// Check errors during iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}

	if len(comments) == 0 {
		return []Comment{}, nil
	}

	return comments, nil
}
