package bookmarks

import (
	"github.com/gandra/my-bookmarks/domain"
	"strings"
)

type BookmarkSearchDto struct {
	domain.Pagination
	Criteria string `json:"criteria"`
}

func (s *BookmarkSearchDto) GetCriteriaLike() string {
	s.Criteria = strings.TrimSpace(s.Criteria)
	if s.Criteria == "" {
		return ""
	}
	return "%" + s.Criteria + "%"
}

type BookmarksDto struct {
	Data []BookmarkDto `json:"data"`
	domain.Pagination
}
