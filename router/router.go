package router

import "github.com/gofiber/fiber/v2"

func SetupRoutes(app *fiber.App) {
	app := fiber.App()
	api := app.Group("/api", logger.New())
}

/* 	Example of grouping
	user := api.Group("user")
	user.GET("/", func(c *fiber.Ctx) {} )
	user.GET("/:userId", func(c *fiber.Ctx) {} )
	user.PUT("/:userId" ,func(c *fiber.Ctx) {} )
*/