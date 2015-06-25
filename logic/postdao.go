package blog

import "time"


type Post struct {
	Title    string
	Content  string
	Author   string
	DateTime time.Time
	Permalink string
}


