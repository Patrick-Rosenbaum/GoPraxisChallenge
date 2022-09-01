package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	// für die Statische Dateien wie CSS, Images und JavaScript
	app.Static("/", "./")

	// gibt Hello World zurück
	app.Get("/todos", func(c *fiber.Ctx) error {
		return c.SendFile("./indexTest.html")
	})

	app.Listen(":5000")
}
