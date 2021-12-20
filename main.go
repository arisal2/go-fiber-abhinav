package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/arisal2/go-fiber-abhinav/database"
)

func main() {
	// Start a new fiber app
	app := fiber.New()

	// Connect to the Databse
	database.ConnectDB()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
}