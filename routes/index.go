package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ngn13/note-server/lib"
)

func GET_Index(c *fiber.Ctx) error {
	notes := c.Locals("notes").([]lib.Note)

	if len(notes) > 20 {
		return c.Render("index", fiber.Map{
			"notes": notes[:20],
		})
	}

	return c.Render("index", fiber.Map{
		"notes": notes,
	})
}
