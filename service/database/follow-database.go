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

func (db *appdbimpl) UnfollowUser(userID int, followID int) error {
	
	_, err := db.c.Exec("DELETE FROM follow WHERE id = ? AND followID = ?", userID, followID)
	if err != nil {
		fmt.Println("DATABASE: Error while following user: ", err)
		return err
	}
	return nil
}

func (db *appdbimpl) GetUserFollowers(userID int) ([]User, error) {

	var followers []User

	query := `SELECT follow.id, users.username
			  FROM follow
			  JOIN users ON users.id = follow.id
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

	return followers, nil
}


func (db *appdbimpl) GetUserFollowing(userID int) ([]User, error) {

	var following []User

	query := `SELECT follow.followID, users.username
			  FROM follow
			  JOIN users ON users.id = follow.followID
			  WHERE follow.id = ?`

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

	return following, nil
}