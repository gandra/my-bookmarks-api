package bookmarks

import (
	"github.com/gandra/my-bookmarks/domain/bookmarks"
	"github.com/gandra/my-bookmarks/logger"
	"github.com/gandra/my-bookmarks/services"
	"github.com/gandra/my-bookmarks/utils/errors"
	"github.com/gandra/my-bookmarks/utils/uri"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func restErrToFiberError(c *fiber.Ctx, err *errors.RestErr) error {
	return c.Status(err.Status).JSON(fiber.Map{"status": "error", "message": err.Message, "data": err.Error})
}

func inputErrorToBadRequestFiberError(c *fiber.Ctx, err error) error {
	logger.Error("input error", err)
	return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Check your input", "data": err})
}

func CreateBookmark(c *fiber.Ctx) error {
	bookmark := new(bookmarks.Bookmark)
	parseErr := c.BodyParser(bookmark)
	if parseErr != nil {
		return inputErrorToBadRequestFiberError(c, parseErr)
	}
	createErr := services.BookmarksService.CreateBookmark(bookmark)
	if createErr != nil {
		return restErrToFiberError(c, createErr)
	}
	return c.JSON(bookmark.Marshall())
}

func GetBookmark(c *fiber.Ctx) error {
	id := c.Params("id")
	bookmark, err := services.BookmarksService.GetBookmark(id)
	if err != nil {
		return restErrToFiberError(c, err)
	}
	return c.JSON(bookmark.Marshall())
}

func SearchBookmarks(c *fiber.Ctx) error {
	criteria := c.Query("criteria")

	searchDto := bookmarks.BookmarkSearchDto{
		Criteria:   criteria,
		Pagination: uri.ParseQueryPagination(c),
	}

	bookmarks, searchErr := services.BookmarksService.SearchBookmarks(searchDto)
	if searchErr != nil {
		return restErrToFiberError(c, searchErr)
	}
	return c.JSON(bookmarks.Marshall())
}
