package routes

import (
	"fiber_blog/app/controllers/web_controllers"
	"fiber_blog/env"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"net/http"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/jhonoryza/inertia-fiber"
	"gorm.io/gorm"
)

func RegisterWebRoute(router *fiber.App, db *gorm.DB) {

	router.Get("/", web_controllers.StaticPage("Home")).Name("home")
	router.Get("/work-with-me", web_controllers.StaticPage("WorkWithMe")).Name("work-with-me")
	router.Get("/disclaimer", web_controllers.StaticPage("Disclaimer")).Name("disclaimer")
	router.Get("/about", web_controllers.StaticPage("About")).Name("about")

	router.Get("/articles", web_controllers.ArticleIndex(db)).Name("articles.index")
	router.Get("/articles/:slug", web_controllers.ArticleShow(db)).Name("articles.show")

	router.Get("/login", web_controllers.LoginForm()).Name("login.form")
	router.Post("/login", web_controllers.Login(db)).Name("login")

	authRouter := router.Group("/auth")

	authRouter.Use(func(c *fiber.Ctx) error {
		cookie := c.Cookies(env.GetEnv().GetString("COOKIE_NAME"))
		if cookie == "" {
			return inertia.Render(c, http.StatusOK, "Error", fiber.Map{
				"message": "Unauthorized",
			})
		}

		c.Locals("user", cookie) // Simpan token di context untuk digunakan di middleware berikutnya
		return c.Next()
	})

	authRouter.Use(jwtware.New(jwtware.Config{
		SigningKey:  jwtware.SigningKey{Key: []byte(env.GetEnv().GetString("JWT_SECRET"))},
		TokenLookup: "cookie:" + env.GetEnv().GetString("COOKIE_NAME"),
		ContextKey:  "user",
	}))

	authRouter.Use(func(c *fiber.Ctx) error {
		user := c.Locals("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)

		name, ok := claims["name"].(string)
		if !ok {
			return c.Next()
		}

		inertia.Share(c, fiber.Map{
			"name": name,
		})

		return c.Next()
	})

	authRouter.Get("dashboard", web_controllers.Dashboard()).Name("dashboard")
	authRouter.Post("logout", web_controllers.Logout()).Name("logout")
	authRouter.Get("profile", web_controllers.Profile()).Name("dashboard")
	authRouter.Get("posts", web_controllers.PostIndex(db)).Name("posts.index")

	fmt.Println("web route register success")
}
