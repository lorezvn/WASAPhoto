package database

func (db *appdbimpl) FollowExists(userID int, followID int) bool {

	var count int
	err := db.c.QueryRow("SELECT COUNT(*) FROM follow WHERE userID = ? AND followID = ?", userID, followID).Scan(&count)
	if err != nil {
		return false
	}

	return count > 0
}

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
		return ErrFollowNotFound
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

	if len(followers) == 0 {
		return []User{}, nil
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

	if len(following) == 0 {
		return []User{}, nil
	}

	return following, nil
}
