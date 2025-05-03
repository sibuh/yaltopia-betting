package main

import (
	"yaltopia-betting/internal/cricket"
	"yaltopia-betting/internal/volleyball"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	//create app instance
	app := fiber.New()
	//add logger middleware
	app.Use(logger.New(logger.Config{

		Format:     "${pid} ${status} - ${method} ${path}\n",
		TimeFormat: "02-Jan-2006",
	}))

	// routes

	// / handler is not part of task just added at
	// the begining to see server started successfully
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
	app.Get("/vb", volleyball.VolleyballBetting)
	app.Get("/crkt", cricket.CricketBetting)

	// server will start at port 3000
	app.Listen(":3000")
}
