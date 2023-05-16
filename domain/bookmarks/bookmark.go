package bookmarks

import "github.com/gandra/my-bookmarks/domain"

type Bookmark struct {
	domain.Base
	Title string `json:"title"`
	Body  string `json:"body"`
	Link  string `json:"link"`
}

type Bookmarks []Bookmark
