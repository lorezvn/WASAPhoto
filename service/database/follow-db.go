package database

import (
	"fmt"
)

func (db *appdbimpl) FollowUser(userID int, followID int) error {
	_, err := db.c.Exec("INSERT INTO follow (id, followID) VALUES (?, ?)", userID, followID)
	if err != nil {
		fmt.Println("DATABASE: Error while following user: ", err)
		return err
	}
	return nil
}
