package noteRoutes

import "github.com/gofiber/fiber/v2"

func SetupNoteRoutes(router fiber.Router) {
	note := router.Group("/note")
	note.Post("/", func(c *fiber.Ctx) error {}) // Create a Note
	note.Get("/", func(c *fiber.Ctx) error {}) // Read all notes
	note.Get("/:noteId", func(c *fiber.Ctx) error {}) // Read one note
	note.Put("/:noteId", func(c *fiber.Ctx) error {}) // Update one note
	note.Delete("/:noteId", func(c *fiber.Ctx) error {}) // Delete one note
}