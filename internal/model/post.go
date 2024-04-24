package model

import (
	"github.com/a-h/templ"
)

type Post struct {
	Title       string   `toml:"title"`
	Author      string   `toml:"author"`
	Description string   `toml:"description"`
	Date        string   `toml:"date"`
	Tags        []string `toml:"tags"`
	Slug        string   `toml:"slug"`
	ContentRaw  string
	Content     templ.Component
}
