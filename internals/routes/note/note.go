package noteRoutes

import ( 
	"github.com/gofiber/fiber/v2"
	noteHandler "github.com/arisal2/go-fiber-abhinav/internals/handlers/note"
)

func SetupNoteRoutes(router fiber.Router) {
	note := router.Group("/note")
	note.Post("/", noteHandler.CreateNotes) // Create a Note
	note.Get("/", noteHandler.GetNotes) // Read all notes
	note.Get("/:noteId", noteHandler.GetNote) // Read one note
}