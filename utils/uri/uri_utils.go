package uri

import (
	"github.com/gandra/my-bookmarks/domain"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"strings"
)

const (
	limitDeafult = 10
	pageDeafult  = 1
	orderDefault = "asc"
)

func ParseQueryPagination(c *fiber.Ctx) domain.Pagination {
	limitStr := strings.TrimSpace(c.Query("limit", ""))
	pageStr := strings.TrimSpace(c.Query("page", ""))
	sort := strings.TrimSpace(c.Query("sort", ""))
	order := strings.ToLower(strings.TrimSpace(c.Query("order", "")))

	limit := limitDeafult
	page := pageDeafult
	if limitStr != "" {
		limit, _ = strconv.Atoi(limitStr)
		if limit < 0 {
			limit = limitDeafult
		}
	}
	if pageStr != "" {
		page, _ = strconv.Atoi(pageStr)
		if page < 1 {
			page = pageDeafult
		}
	}
	if order != "asc" && order != "desc" {
		order = orderDefault
	}

	return domain.Pagination{
		Limit: limit,
		Page:  page,
		Sort:  sort,
		Order: order,
	}
}
