package model

import (
	"time"

	"github.com/a-h/templ"
)

type Post struct {
	Title       string    `toml:"title"`
	Author      string    `toml:"author"`
	Description string    `toml:"description"`
	Date        time.Time `toml:"date"`
	Tags        []string  `toml:"tags"`
	Slug        string    `toml:"slug"`
	Content     templ.Component
}
