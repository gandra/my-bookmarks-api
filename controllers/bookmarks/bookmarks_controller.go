package bookmarks

import (
	"github.com/gandra/my-bookmarks/domain/bookmarks"
	"github.com/gandra/my-bookmarks/services"
	"github.com/gandra/my-bookmarks/utils/errors"
	"github.com/gofiber/fiber/v2"
)

func getFiberError(c *fiber.Ctx, err *errors.RestErr) error {
	return c.Status(err.Status).JSON(fiber.Map{"status": "error", "message": err.Message, "data": err.Error})
}

func GetBookmark(c *fiber.Ctx) error {
	id := c.Params("id")
	bookmark, err := services.BookmarksService.GetBookmark(id)
	if err != nil {
		return getFiberError(c, err)
	}
	return c.JSON(bookmark)
}

func SearchBookmarks(c *fiber.Ctx) error {
	criteria := c.Params("criteria")
	criteriaDto := bookmarks.BookmarkSearchDto{
		Criteria: criteria,
	}
	bookmarks, err := services.BookmarksService.SearchBookmarks(criteriaDto)
	if err != nil {
		return getFiberError(c, err)
	}
	return c.JSON(bookmarks)
}
