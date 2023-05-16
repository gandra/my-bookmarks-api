package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func StartApplication() {
	app := fiber.New()
	app.Use(cors.New())
	SetupRoutes(app)
	app.Listen(":8000")
}
