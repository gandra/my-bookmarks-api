package repository

import (
	"github.com/gandra/my-bookmarks/database/db"
	"github.com/gandra/my-bookmarks/domain/bookmarks"
	"strings"
)

var (
	BookmarksRepository bookmarksRepositoryInterface = &bookmarksRepository{}
)

type bookmarksRepository struct {
}

type bookmarksRepositoryInterface interface {
	GetById(string) (*bookmarks.Bookmark, error)
	Create(*bookmarks.Bookmark) error
	Search(bookmarks.BookmarkSearchDto) (*bookmarks.Bookmarks, error)
}

func (r *bookmarksRepository) GetById(id string) (*bookmarks.Bookmark, error) {
	var bookmark bookmarks.Bookmark

	err := db.Client.Find(&bookmark, id).Error
	if err != nil {
		return nil, err
	}
	return &bookmark, nil
}

func (r *bookmarksRepository) Create(bookmark *bookmarks.Bookmark) error {
	return db.Client.Create(bookmark).Error
}

func (r *bookmarksRepository) Search(searchDto bookmarks.BookmarkSearchDto) (*bookmarks.Bookmarks, error) {
	var bookmarks bookmarks.Bookmarks

	fields := make([]string, 0)
	values := make([]interface{}, 0)
	if searchDto.GetCriteriaLike() != "" {
		fields = append(fields, "lower(title) like ?")
		values = append(values, strings.ToLower(searchDto.GetCriteriaLike()))
	}

	pagination := searchDto.Pagination
	bookmarks.Pagination = pagination
	offset := (pagination.Page - 1) * pagination.Limit
	queryBuider := db.Client.Limit(pagination.Limit).Offset(offset).Order(pagination.GetSortOrder())
	err := queryBuider.Where(strings.Join(fields, " AND "), values...).Find(&bookmarks.Data).Error
	//err := db.Client.Where(strings.Join(fields, " AND "), values...).Find(&bookmarks.Data).Error
	if err != nil {
		return nil, err
	}
	return &bookmarks, nil
}
