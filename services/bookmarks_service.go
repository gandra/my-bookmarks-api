package services

import (
	"github.com/gandra/my-bookmarks/domain/bookmarks"
	"github.com/gandra/my-bookmarks/logger"
	"github.com/gandra/my-bookmarks/repository"
	"github.com/gandra/my-bookmarks/utils/errors"
)

var (
	BookmarksService bookmarksServiceInterface = &bookmarksService{}
)

type bookmarksServiceInterface interface {
	CreateBookmark(*bookmarks.Bookmark) *errors.RestErr
	GetBookmark(string) (*bookmarks.Bookmark, *errors.RestErr)
	SearchBookmarks(bookmarks.BookmarkSearchDto) (*bookmarks.Bookmarks, *errors.RestErr)
}
type bookmarksService struct {
}

func (s *bookmarksService) GetBookmark(id string) (*bookmarks.Bookmark, *errors.RestErr) {
	bookmark, err := repository.BookmarksRepository.GetById(id)
	if err != nil {
		logger.Error("Error finding bookmark by id", err)
		return nil, errors.NewInternalServerError(err.Error())
	}
	return bookmark, nil
}

func (s *bookmarksService) CreateBookmark(bookmark *bookmarks.Bookmark) *errors.RestErr {
	err := repository.BookmarksRepository.Create(bookmark)
	if err != nil {
		logger.Error("Error creating bookmark", err)
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}

func (s *bookmarksService) SearchBookmarks(searchDto bookmarks.BookmarkSearchDto) (*bookmarks.Bookmarks, *errors.RestErr) {
	bookmarks, err := repository.BookmarksRepository.Search(searchDto)
	if err != nil {
		logger.Error("Error searching bookmarks", err)
		return nil, errors.NewInternalServerError(err.Error())
	}
	return bookmarks, nil
}
