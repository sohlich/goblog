package repository

import(
	"time"
)

type Post struct {
	Title    string
	Content  string
	Author   string
	DateTime time.Time
	Permalink string
	Tags	[]string
}


type User struct{
	Username string
	Password string
	Email string
	SessionToken string
}