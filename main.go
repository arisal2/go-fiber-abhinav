package main

import (
	"encoding/json"
	"github.com/arisal2/go-fiber-abhinav/database"
	"github.com/arisal2/go-fiber-abhinav/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Start a new fiber app
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	// Connect to the Database
	database.ConnectDB()

	// Setup the router
	router.SetupRoutes(app)

	// Listen on PORT 3000
	app.Listen(":3000")
}
