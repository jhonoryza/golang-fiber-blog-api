package web_controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/jhonoryza/inertia-fiber"
)

func Dashboard() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return inertia.Render(c, http.StatusOK, "Admin/Dashboard", fiber.Map{})
	}
}
