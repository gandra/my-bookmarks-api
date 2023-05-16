package app

import (
	"github.com/gandra/my-bookmarks/controllers/bookmarks"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/api/bookmarks", bookmarks.SearchBookmarks)
	app.Get("/api/bookmarks/:id", bookmarks.GetBookmark)
}
