package database

import (
	"errors"
)

func (db *appdbimpl) BanUser(userID int, banID int) error {

	/*
	var count int 
	err := db.c.QueryRow("SELECT COUNT(*) FROM ban WHERE userID = ? AND banID = ?", userID, banID).Scan(&count)
	if err != nil {
		return err
	}

	if count > 0 {
		return errors.New("User already banned")
	}
	*/

	_, err := db.c.Exec("INSERT INTO ban (userID, banID) VALUES (?, ?)", userID, banID)
	if err != nil {
		return err
	}

	_ = db.UnfollowUser(userID, banID) 
	return nil
}

func (db *appdbimpl) UnbanUser(userID int, banID int) error {
	
	result, err := db.c.Exec("DELETE FROM ban WHERE userID = ? AND banID = ?", userID, banID)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("Ban not found")
	}
	return nil
}