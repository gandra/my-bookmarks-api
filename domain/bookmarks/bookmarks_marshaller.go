package bookmarks

import (
	"github.com/gandra/my-bookmarks/utils/date_utils"
)

type BookmarkDto struct {
	Id        string `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	Link      string `json:"link"`
}

func (b Bookmarks) Marshall() interface{} {
	data := make([]BookmarkDto, len(b.Data))
	for i, bookmark := range b.Data {
		data[i] = bookmark.Marshall().(BookmarkDto)
	}
	return BookmarksDto{
		Data:       data,
		Pagination: b.Pagination,
	}
}

func (b *Bookmark) Marshall() interface{} {
	return BookmarkDto{
		Id:        b.Id.String(),
		CreatedAt: date_utils.GetDateString(b.CreatedAt),
		UpdatedAt: date_utils.GetDateString(b.UpdatedAt),
		Title:     b.Title,
		Body:      b.Body,
		Link:      b.Link,
	}
}
