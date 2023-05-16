package app

import (
	"github.com/gandra/my-bookmarks/database/conf_db"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func StartApplication() {
	conf_db.InitDatabase()
	app := fiber.New()
	app.Use(cors.New())
	SetupRoutes(app)
	app.Listen(":8000")
}
