package noteRoutes

import (
	noteHandler "github.com/arisal2/go-fiber-abhinav/internals/handlers/note"
	"github.com/gofiber/fiber/v2"
)

func SetupNoteRoutes(router fiber.Router) {
	note := router.Group("/note")
	note.Post("/", noteHandler.CreateNotes)
	note.Get("/", noteHandler.GetNotes)
	note.Get("/:noteId", noteHandler.GetNote)
	note.Put("/:noteId", noteHandler.UpdateNote)
	note.Delete("/:noteId", noteHandler.DeleteNote)
}
