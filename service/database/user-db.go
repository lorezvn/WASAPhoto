package database

import (
	"database/sql"
	"errors"
	"fmt"
)

func (db *appdbimpl) GetUserByUsername(username string) (int, error) {
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

func (db *appdbimpl) GetUserById(userID int) (int, error) {

	var id int
	err := db.c.QueryRow("SELECT id FROM users WHERE id = ?", userID).Scan(&id)
	if err != nil {
		return 0, err
	}
	return 0, nil
}

func (db *appdbimpl) ChangeUsername(userID int, newUsername string) error {

	_, err := db.GetUserById(userID)
	if err != nil {
		fmt.Print("DATABASE: Error obtaining user id: ", err)
		return err
	}

	_, err = db.c.Exec("UPDATE users SET username = ? WHERE id = ?", newUsername, userID)
	if err != nil {
		fmt.Println("DATABASE: Error updating username: ", err)
		return err
	}

	return nil
}

func (db *appdbimpl) CreateUser(username string) (int, error) {

	userID, err := db.GetUserByUsername(username)
	if err != nil {
		fmt.Println("DATABASE: Error getting the id of the User by username: ", err)
		return 0, err
	}

	// If User already exists -> Return the identifier
	if userID != 0 {
		return userID, nil
	}

	// Else create a new User and return the new identifier
	result, err := db.c.Exec("INSERT INTO users (username) VALUES (?)", username)
	if err != nil {
		fmt.Println("DATABASE: Error inserting user: ", err)
		return 0, err
	}

	newUserID, err := result.LastInsertId()
	if err != nil {
		fmt.Println("DATABASE: Error obtaining the last insert id: ", err)
		return 0, err
	}

	return int(newUserID), nil
}
