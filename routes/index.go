package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ngn13/note-server/lib"
)

func GetIndex(c *fiber.Ctx) error {
  var first []lib.Note 
  all := lib.GetNotes()

  for _, n := range all {
    if (len(first) == 20) {
      break
    }
    first = append(first, n)
  }

  return c.Render("index", fiber.Map{
    "notes": first, 
  })
}
