package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/django/v3"
	"github.com/ngn13/note-server/lib"
	"github.com/ngn13/note-server/routes"
)

func main(){
  lib.CheckNotes()

  engine := django.New("./views", ".html")
  app := fiber.New(fiber.Config{
    Views: engine,
  })

  addr := os.Getenv("ADDR")
  if addr == "" {
    addr = ":8080"
  }
  
  app.Use(logger.New())
  app.Static("/", "./static")

  app.Get("/", routes.GetIndex)
  app.Get("/search", routes.GetSearch)
  app.Get("*", routes.GetNote)

  log.Infof("Starting application on %s", addr)
  log.Fatal(app.Listen(addr))
}
