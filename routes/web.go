package routes

import (
	"fiber_blog/app/controllers/web_controllers"
	"fmt"
	"github.com/gofiber/fiber/v2"

	"gorm.io/gorm"
)

func RegisterWebRoute(router *fiber.App, db *gorm.DB) {

	router.Get("/", web_controllers.StaticPage("Home")).Name("home")
	router.Get("/work-with-me", web_controllers.StaticPage("WorkWithMe")).Name("work-with-me")
	router.Get("/disclaimer", web_controllers.StaticPage("Disclaimer")).Name("disclaimer")
	router.Get("/about", web_controllers.StaticPage("About")).Name("about")

	router.Get("/articles", web_controllers.ArticleIndex(db)).Name("articles.index")
	router.Get("/articles/:slug", web_controllers.ArticleShow(db)).Name("articles.show")

	fmt.Println("web route register success")
}
