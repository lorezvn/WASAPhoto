package api

type User struct {
	UserID   int    `json:"userID"`
	Username string `json:"username"`
}

type UserProfile struct {
	UserID    int     `json:"userID"`
	Username  string  `json:"username"`
	Followers []User  `json:"followers"`
	Following []User  `json:"following"`
	Photos    []Photo `json:"photos"`
}

type Photo struct {
	Image    []byte    `json:"image"`
	PhotoID  int       `json:"photoID"`
	UserID   int       `json:"userID"`
	Date     string    `json:"date"`
	Likes    []Like    `json:"likes"`
	Comments []Comment `json:"comments"`
}

type Comment struct {
	CommentID int    `json:"commentID"`
	UserID    int    `json:"userID"`
	PhotoID   int    `json:"photoID"`
	Message   string `json:"message"`
	Date      string `json:"date"`
}

type Like struct {
	UserID  int    `json:"userID"`
	PhotoID int    `json:"photoID"`
	Date    string `json:"date"`
}
