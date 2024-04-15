package model

import "time"

type Post struct {
	Title       string
	Author      string
	Description string
	Date        time.Time
	Tags        []string
	Slug        string
	Content     string
}
