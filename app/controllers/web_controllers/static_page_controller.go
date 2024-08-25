package web_controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jhonoryza/inertia-fiber"
)

func StaticPage(template string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return inertia.Render(c, 200, template, fiber.Map{})
	}
}
