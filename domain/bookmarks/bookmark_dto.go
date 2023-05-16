package bookmarks

import "strings"

type BookmarkSearchDto struct {
	Criteria string
}

func (s *BookmarkSearchDto) GetCriteriaLike() string {
	s.Criteria = strings.TrimSpace(s.Criteria)
	if s.Criteria == "" {
		return ""
	}
	return "%" + s.Criteria + "%"
}
