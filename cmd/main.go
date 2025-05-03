package main

import (
	"yaltopia-betting/internal/cricket"
	"yaltopia-betting/internal/volleyball"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(logger.New(logger.Config{

		Format:     "${pid} ${status} - ${method} ${path}\n",
		TimeFormat: "02-Jan-2006",
	}))
	// ... (define routes and other settings) ...
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
	app.Get("/vb", volleyball.VolleyballBetting)
	app.Get("/cricket", cricket.CricketBetting)

	app.Get("/users/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		return c.SendString("User ID: " + id)
	})
	app.Listen(":3000") // Start the server on port 3000
}
