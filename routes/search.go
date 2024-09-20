package routes

import (
	"bytes"
	"strings"

	"github.com/gofiber/fiber/v2"
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

func GET_Search(c *fiber.Ctx) error {
	var (
		search  string
		content []byte
		notes   []lib.Note
		results []lib.Note
		indexes []int
		err     error
	)

	search = lib.XSS_sanitizer.Sanitize(c.Query("s"))
	notes = c.Locals("notes").([]lib.Note)

	if search == "" {
		return c.Render("index", fiber.Map{
			"search": search,
			"notes":  notes,
		})
	}

	for i, n := range notes {
		if strings.Contains(n.Path, search) {
			results = append(results, n)
			indexes = append(indexes, i)
		}
	}

	for i, n := range notes {
		if content, err = n.Read(); err != nil {
			lib.Fail("failed to read the note \"%s\" during search: %s", err.Error())
		}

		if !contains(indexes, i) && bytes.Contains(content, []byte(search)) {
			results = append(results, n)
		}
	}

	return c.Render("index", fiber.Map{
		"search": search,
		"notes":  results,
	})
}
