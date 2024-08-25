package routes

import (
	"fiber_blog/app/controllers/api_controllers"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterApiRoute(router *fiber.App, db *gorm.DB) {
	api := router.Group("/api")

	api.Get("/posts", api_controllers.PostIndex(db))
	api.Get("/posts/:slug", api_controllers.PostShow(db))

	fmt.Println("api route register success")
}
