package web_controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jhonoryza/inertia-fiber"
)

func Dashboard() fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.Locals("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)

		name := claims["name"].(string)

		return inertia.Render(c, http.StatusOK, "Admin/Dashboard", fiber.Map{
			"name": name,
		})
	}
}
