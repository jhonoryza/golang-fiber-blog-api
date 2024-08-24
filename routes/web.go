package routes

import (
	"fiber_blog/app/controllers"
	"fiber_blog/config"
	"fmt"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/helmet"
)

func RegisterWebRoute() {
	router := config.GetRouter()

	web := router.Group("/")

	web.Use(csrf.New(csrf.Config{
		KeyLookup:      "header:X-XSRF-TOKEN",
		CookieName:     "XSRF-TOKEN",
		SingleUseToken: true,
	}))

	web.Use(helmet.New(helmet.Config{
		ReferrerPolicy: "strict-origin-when-cross-origin",
	}))

	web.Get("/", controllers.HomeIndex)

	fmt.Println("web route register success")
}
