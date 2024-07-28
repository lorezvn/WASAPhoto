/*
Package database is the middleware between the app database and the code. All data (de)serialization (save/load) from a
persistent database are handled here. Database specific logic should never escape this package.

To use this package you need to apply migrations to the database if needed/wanted, connect to it (using the database
data source name from config), and then initialize an instance of AppDatabase from the DB connection.

For example, this code adds a parameter in `webapi` executable for the database data source name (add it to the
main.WebAPIConfiguration structure):

	DB struct {
		Filename string `conf:""`
	}

This is an example on how to migrate the DB and connect to it:

	// Start Database
	logger.Println("initializing database support")
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		logger.WithError(err).Error("error opening SQLite DB")
		return fmt.Errorf("opening SQLite: %w", err)
	}
	defer func() {
		logger.Debug("database stopping")
		_ = db.Close()
	}()

Then you can initialize the AppDatabase and pass it to the api package.
*/
package database

import (
	"database/sql"
	"errors"
	"fmt"
)

// AppDatabase is the high level interface for the DB
type AppDatabase interface {
	GetName() (string, error)
	SetName(name string) error

	UserExists(userID int) bool
	UsernameExists(username string) bool
	PhotoExists(photoID int, photoAuthorID int) bool
	BanExists(userID int, banID int) bool

	GetUsernameByUserID(userID int) (string, error)
	CreateUser(username string) (int, error)
	ChangeUsername(userID int, newUsername string) error
	SearchUsers(username string) ([]User, error)
	InsertPhoto(userID int, image []byte) (int, string, error)
	DeletePhoto(photoID int) error
	InsertComment(userID int, photoID int, message string) (int, string, error)
	DeleteComment(commentID int) error
	InsertLike(userID int, photoID int) (string, error)
	DeleteLike(userID int, photoID int) error
	FollowUser(userID int, followID int) error
	UnfollowUser(userID int, followID int) error
	BanUser(userID int, banID int) error
	UnbanUser(userID int, banID int) error
	GetStream(userID int) ([]Photo, error)
	GetUserProfile(userID int) (UserProfile, error)

	Ping() error
}

type appdbimpl struct {
	c *sql.DB
}

// New returns a new instance of AppDatabase based on the SQLite connection `db`.
// `db` is required - an error will be returned if `db` is `nil`.
func New(db *sql.DB) (AppDatabase, error) {
	if db == nil {
		return nil, errors.New("database is required when building a AppDatabase")
	}

	// Check if table exists. If not, the database is empty, and we need to create the structure
	var tableName string
	err := db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='example_table';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `PRAGMA foreign_keys=on;
					CREATE TABLE IF NOT EXISTS users (
						id INTEGER PRIMARY KEY AUTOINCREMENT, 
						username TEXT
					);
					CREATE TABLE IF NOT EXISTS photos (
						id INTEGER PRIMARY KEY AUTOINCREMENT,
						userID INTEGER,
						image BLOB,
						date TEXT,
						FOREIGN KEY (userID) REFERENCES users(id)
					);
					CREATE TABLE IF NOT EXISTS comments (
						id INTEGER PRIMARY KEY AUTOINCREMENT,
						userID INTEGER,
						photoID INTEGER,
						message TEXT,
						date TEXT,
						FOREIGN KEY (userID) REFERENCES users(id)
						FOREIGN KEY (photoID) REFERENCES photos(id)
					);
					CREATE TABLE IF NOT EXISTS likes (
						userID INTEGER,
						photoID INTEGER,
						date TEXT,
						PRIMARY KEY (userID, photoID)
						FOREIGN KEY (userID) REFERENCES users(id)
						FOREIGN KEY (photoID) REFERENCES photos(id)
					);
					CREATE TABLE IF NOT EXISTS follow (
						userID INTEGER, 
						followID INTEGER,
						PRIMARY KEY (userID, followID),
						FOREIGN KEY (userID) REFERENCES users(id),
						FOREIGN KEY (followID) REFERENCES users(id)
					);
					CREATE TABLE IF NOT EXISTS ban (
						userID INTEGER, 
						banID INTEGER,
						PRIMARY KEY (userID, banID),
						FOREIGN KEY (userID) REFERENCES users(id),
						FOREIGN KEY (banID) REFERENCES users(id)
					);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	return &appdbimpl{
		c: db,
	}, nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
