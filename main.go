package main

import (
	"flag"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/django/v3"
	"github.com/ngn13/note-server/lib"
	"github.com/ngn13/note-server/routes"
)

func main() {
	var (
		opt_interface *string
		opt_staticdir *string
		opt_viewsdir  *string
		opt_notesdir  *string
		err           error
	)

	opt_interface = flag.String("interface", "127.0.0.1:8080", "Web server interface in host:port format")
	opt_staticdir = flag.String("static", "/usr/lib/note-server/static", "Static files directory path")
	opt_viewsdir = flag.String("views", "/usr/lib/note-server/views", "HTML templates directory path")
	opt_notesdir = flag.String("notes", "", "Path for the directory that contains your notes")
	flag.Parse()

	if *opt_notesdir == "" {
		lib.Info("please specify a notes directory using the -notes option")
		os.Exit(0)
	}

	if _, err = os.Stat(*opt_notesdir); err != nil {
		lib.Fail("cannot access to notes directory: %s", *opt_notesdir)
		os.Exit(1)
	}

	if _, err = os.Stat(*opt_staticdir); err != nil {
		lib.Fail("cannot access to static directory: %s", *opt_staticdir)
		os.Exit(1)
	}

	if _, err = os.Stat(*opt_viewsdir); err != nil {
		lib.Fail("cannot access to views directory: %s", *opt_viewsdir)
		os.Exit(1)
	}

	if *opt_interface == "" {
		lib.Fail("invalid web server interface: %s", *opt_interface)
		os.Exit(1)
	}

	engine := django.New(*opt_viewsdir, ".html")

	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		Views:                 engine,
	})

	app.All("*", func(c *fiber.Ctx) error {
		var notes []lib.Note

		if notes, err = lib.GetNotes(*opt_notesdir); err != nil {
			lib.Fail("cannot obtain notes: %s", err.Error())
			return c.Status(404).Render("error", fiber.Map{
				"message": "500 Internal Server Error",
			})
		}

		c.Locals("notes", notes)
		c.Locals("dir", *opt_notesdir)
		return c.Next()
	})

	app.Static("/", *opt_staticdir)

	app.Get("/", routes.GET_Index)
	app.Get("/search", routes.GET_Search)
	app.Get("*", routes.GET_note)

	lib.Info("starting application on %s", *opt_interface)

	if err = app.Listen(*opt_interface); err != nil {
		lib.Fail("cannot start the application: %s", err.Error())
		os.Exit(1)
	}
}
