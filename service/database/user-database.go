package database

import (
	"database/sql"
	"errors"
)

func (db *appdbimpl) GetUserIDByUsername(username string) (int, error) {
	var userID int
	err := db.c.QueryRow("SELECT id FROM users WHERE username = ?", username).Scan(&userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, nil
		}
		return 0, err
	}

	return userID, nil
}

func (db *appdbimpl) GetUsernameByUserID(userID int) (string, error) {
	var username string
	err := db.c.QueryRow("SELECT username FROM users WHERE id = ?", userID).Scan(&username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", nil
		}
		return "", err
	}

	return username, nil
}

func (db *appdbimpl) SearchUsers(username string) ([]User, error) {

	var users []User

	query := `SELECT * FROM users WHERE username LIKE ?`

	rows, err := db.c.Query(query, username+"%")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {

		var user User
		if err := rows.Scan(&user.UserID, &user.Username); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	// Check errors during iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return []User{}, nil
	}

	return users, nil
}

func (db *appdbimpl) UserExists(userID int) bool {

	var id int
	err := db.c.QueryRow("SELECT id FROM users WHERE id = ?", userID).Scan(&id)
	return err == nil
}

func (db *appdbimpl) UsernameExists(userID int, username string) bool {

	var count int

	// Excluding the user making the request, this allows the user to change to the username they previously had
	err := db.c.QueryRow("SELECT COUNT(*) FROM users WHERE username = ? AND id <> ?", username, userID).Scan(&count)
	if err != nil {
		return false
	}
	return count > 0
}

func (db *appdbimpl) ChangeUsername(userID int, newUsername string) error {

	_, err := db.c.Exec("UPDATE users SET username = ? WHERE id = ?", newUsername, userID)
	if err != nil {
		return err
	}

	return nil
}

func (db *appdbimpl) CreateUser(username string) (int, error) {

	userID, err := db.GetUserIDByUsername(username)
	if err != nil {
		return 0, err
	}

	// If User already exists -> Return the identifier
	if userID != 0 {
		return userID, nil
	}

	// Else create a new User and return the new identifier
	result, err := db.c.Exec("INSERT INTO users (username) VALUES (?)", username)
	if err != nil {
		return 0, err
	}

	newUserID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(newUserID), nil
}

func (db *appdbimpl) GetUserProfile(userID int) (UserProfile, error) {

	userProfile := UserProfile{
		UserID: userID,
	}

	var username string

	err := db.c.QueryRow("SELECT username FROM users WHERE id = ?", userID).Scan(&username)
	if err != nil {
		return userProfile, err
	}

	userProfile.Username = username

	userProfile.Photos, err = db.GetUserPhotos(userID)
	if err != nil {
		return userProfile, err
	}

	userProfile.Following, err = db.GetUserFollowing(userID)
	if err != nil {
		return userProfile, err
	}

	userProfile.Followers, err = db.GetUserFollowers(userID)
	if err != nil {
		return userProfile, err
	}

	return userProfile, nil
}
