package services

import (
	"github.com/gandra/my-bookmarks/domain/bookmarks"
	"github.com/gandra/my-bookmarks/utils/errors"
)

var (
	BookmarksService bookmarksServiceInterface = &bookmarksService{}
)

type bookmarksServiceInterface interface {
	GetBookmark(string) (*bookmarks.Bookmark, *errors.RestErr)
	SearchBookmarks(bookmarks.BookmarkSearchDto) (bookmarks.Bookmarks, *errors.RestErr)
}
type bookmarksService struct {
}

func (s *bookmarksService) GetBookmark(id string) (*bookmarks.Bookmark, *errors.RestErr) {
	return &bookmarks.Bookmark{
		Id:        "1",
		Title:     "Some Bookmark",
		Body:      "Lorem ipsum dolor ...",
		Link:      "https://www.google.com/",
		CreatedAt: "",
		UpdatedAt: "",
	}, nil
}

func (s *bookmarksService) SearchBookmarks(searchDto bookmarks.BookmarkSearchDto) (bookmarks.Bookmarks, *errors.RestErr) {
	result := make([]bookmarks.Bookmark, 0)
	result = append(result, bookmarks.Bookmark{
		Id:        "1",
		Title:     "Some Bookmark",
		Body:      "Lorem ipsum dolor ...",
		Link:      "https://www.google.com/",
		CreatedAt: "",
		UpdatedAt: "",
	})
	return result, nil
}
