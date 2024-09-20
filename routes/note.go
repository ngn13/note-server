package routes

import (
	"path"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/ngn13/note-server/lib"
	"github.com/russross/blackfriday/v2"
)

func GET_note(c *fiber.Ctx) error {
	var (
		notes   []lib.Note
		content []byte
		name    string
		dir     string
		fp      string
		op      string
		err     error
	)

	notes = c.Locals("notes").([]lib.Note)
	dir = c.Locals("dir").(string)

	fp = c.Path()
	if fp[0] == '/' {
		fp = fp[1:]
	}

	op = lib.XSS_sanitizer.Sanitize(fp)

	for _, f := range lib.LFI_filters {
		if strings.Contains(fp, f) {
			return GET_NotFound(c)
		}
	}

	if fp == "/" {
		return GET_NotFound(c)
	}

	if !strings.HasSuffix(fp, ".md") {
		_, name = path.Split(fp)

		if fp = lib.FindFile(name, dir); fp == "" {
			lib.Fail("requested file \"%s\" not found", name)
			return GET_NotFound(c)
		}

		return c.SendFile(fp)
	}

	fp = path.Clean(fp)

	for _, n := range notes {
		if n.Name != fp {
			continue
		}

		if content, err = n.Read(); err != nil {
			return c.Status(404).Render("error", fiber.Map{
				"message": "500 Internal Server Error",
			})
		}

		ext := blackfriday.FencedCode
		ext |= blackfriday.SpaceHeadings
		ext |= blackfriday.NoEmptyLineBeforeBlock

		md := blackfriday.Run(
			content,
			blackfriday.WithExtensions(ext),
		)

		return c.Render("note", fiber.Map{
			"path":     op,
			"markdown": string(md),
		})
	}

	lib.Fail("requested note \"%s\" not found", fp)
	return GET_NotFound(c)
}
