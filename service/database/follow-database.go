package database

import (
	"errors"
)

func (db *appdbimpl) FollowUser(userID int, followID int) error {

	_, err := db.c.Exec("INSERT INTO follow (userID, followID) VALUES (?, ?)", userID, followID)
	if err != nil {
		return err
	}
	return nil
}

func (db *appdbimpl) UnfollowUser(userID int, followID int) error {

	result, err := db.c.Exec("DELETE FROM follow WHERE userID = ? AND followID = ?", userID, followID)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("Follow not found")
	}
	return nil
}

func (db *appdbimpl) GetUserFollowers(userID int) ([]User, error) {

	var followers []User

	query := `SELECT follow.userID, users.username
			  FROM follow
			  JOIN users ON users.id = follow.userID
			  WHERE follow.followID = ?`

	rows, err := db.c.Query(query, userID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {

		var user User
		if err := rows.Scan(&user.UserID, &user.Username); err != nil {
			return nil, err
		}

		followers = append(followers, user)
	}

	// Check errors during iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return followers, nil
}

func (db *appdbimpl) GetUserFollowing(userID int) ([]User, error) {

	var following []User

	query := `SELECT follow.followID, users.username
			  FROM follow
			  JOIN users ON users.id = follow.followID
			  WHERE follow.userID = ?`

	rows, err := db.c.Query(query, userID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {

		var user User
		if err := rows.Scan(&user.UserID, &user.Username); err != nil {
			return nil, err
		}

		following = append(following, user)
	}

	// Check errors during iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return following, nil
}
