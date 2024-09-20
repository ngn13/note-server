package routes

import "github.com/gofiber/fiber/v2"

func GET_NotFound(c *fiber.Ctx) error {
	return c.Status(404).Render("error", fiber.Map{
		"message": "404 Not Found",
	})
}
