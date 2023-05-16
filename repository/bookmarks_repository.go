package repository

import (
	"github.com/gandra/my-bookmarks/database/db"
	"github.com/gandra/my-bookmarks/domain/bookmarks"
	"strings"
)

const (
	querySearchBookmarks = "SELECT id, title, body, link, created_at, updated_at where title"
)

var (
	BookmarksRepository bookmarksRepositoryInterface = &bookmarksRepository{}
)

type bookmarksRepository struct {
	//mutex sync.Mutex
}

type bookmarksRepositoryInterface interface {
	GetById(string) (*bookmarks.Bookmark, error)
	Create(*bookmarks.Bookmark) error
	Search(bookmarks.BookmarkSearchDto) (*bookmarks.Bookmarks, error)
}

func (r *bookmarksRepository) GetById(id string) (*bookmarks.Bookmark, error) {
	//r.mutex.Lock()
	//defer r.mutex.Unlock()

	var bookmark bookmarks.Bookmark

	err := db.Client.Find(&bookmark, id).Error
	if err != nil {
		return nil, err
	}
	return &bookmark, nil
}

func (r *bookmarksRepository) Create(bookmark *bookmarks.Bookmark) error {
	//r.mutex.Lock()
	//defer r.mutex.Unlock()

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
	err := db.Client.Where(strings.Join(fields, " AND "), values...).Find(&bookmarks).Error
	//err := db.Client.Where("title like ?", searchDto.GetCriteriaLike()).Find(&bookmarks).Error
	if err != nil {
		return nil, err
	}
	return &bookmarks, nil
}
