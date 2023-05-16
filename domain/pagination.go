package domain

import "fmt"

type Pagination struct {
	Limit int    `json:"limit"`
	Page  int    `json:"page"`
	Sort  string `json:"sort"`
	Order string `json:"order"`
}

func (p *Pagination) GetSortOrder() string {
	if p.Sort == "" {
		return ""
	}
	return fmt.Sprintf("%s %s", p.Sort, p.Order)
}
