package routes

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/microcosm-cc/bluemonday"
	"github.com/ngn13/note-server/lib"
)

func contains(s []int, e int) bool {
  for _, a := range s { 
    if a == e {
      return true 
    }
  }
  return false
}

func GetSearch(c *fiber.Ctx) error {
  search := c.Query("s")
  all := lib.GetNotes()

  // no you aint getting any xss
  s := bluemonday.UGCPolicy()
  search = s.Sanitize(search)

  if search == "" {
    return c.Render("index", fiber.Map{
      "search": search,
      "notes": all, 
    })
  }

  var results []lib.Note
  var indxs []int

  for i, n := range all {
    if strings.Contains(n.Path, search) {
      indxs = append(indxs, i)
      results = append(results, n) 
    }
  }

  for i, n := range all {
    if !contains(indxs, i) && strings.Contains(n.Content, search) {
      results = append(results, n) 
    }
  }

  return c.Render("index", fiber.Map{
    "search": search,
    "notes": results, 
  })
}
