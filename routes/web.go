package routes

import (
	"fiber_blog/app/controllers/web_controllers"
	"fiber_blog/env"
	"fmt"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jhonoryza/inertia-fiber"
	"gorm.io/gorm"
)

func RegisterWebRoute(router *fiber.App, db *gorm.DB) {

	// inertia middleware to share common data
	router.Use(
		jwtware.New(jwtware.Config{
			SigningKey:  jwtware.SigningKey{Key: []byte(env.GetEnv().GetString("JWT_SECRET"))},
			TokenLookup: "cookie:" + env.GetEnv().GetString("COOKIE_NAME"),
			ContextKey:  "user",
			ErrorHandler: func(ctx *fiber.Ctx, err error) error {
				return ctx.Next()
			},
			SuccessHandler: func(ctx *fiber.Ctx) error {
				user := ctx.Locals("user").(*jwt.Token)
				claims := user.Claims.(jwt.MapClaims)

				name, ok := claims["name"].(string)
				if !ok {
					return ctx.Next()
				}

				inertia.Share(ctx, fiber.Map{
					"name": name,
				})

				return ctx.Next()
			},
		}),
	)

	router.Get("/", web_controllers.StaticPage("Home")).Name("home")
	router.Get("/work-with-me", web_controllers.StaticPage("WorkWithMe")).Name("work-with-me")
	router.Get("/disclaimer", web_controllers.StaticPage("Disclaimer")).Name("disclaimer")
	router.Get("/about", web_controllers.StaticPage("About")).Name("about")

	router.Get("/articles", web_controllers.ArticleIndex(db)).Name("articles.index")
	router.Get("/articles/:slug", web_controllers.ArticleShow(db)).Name("articles.show")

	router.Get("/login", web_controllers.LoginForm()).Name("login.form")
	router.Post("/login", web_controllers.Login(db)).Name("login")

	authRouter := router.Group("/auth").Name("auth.")

	// redirect to login form if cookie is empty
	authRouter.Use(func(c *fiber.Ctx) error {
		cookie := c.Cookies(env.GetEnv().GetString("COOKIE_NAME"))
		if cookie == "" {
			return inertia.RedirectToRoute(c, "login.form", fiber.Map{})
		}

		// save token in context tobe used in the nex middleware
		c.Locals("user", cookie)
		return c.Next()
	})

	// redirect to login form if not authenticated
	authRouter.Use(jwtware.New(jwtware.Config{
		SigningKey:  jwtware.SigningKey{Key: []byte(env.GetEnv().GetString("JWT_SECRET"))},
		TokenLookup: "cookie:" + env.GetEnv().GetString("COOKIE_NAME"),
		ContextKey:  "user",
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return inertia.RedirectToRoute(c, "login.form", fiber.Map{})
		},
	}))

	authRouter.Get("dashboard", web_controllers.Dashboard()).Name("dashboard")
	authRouter.Post("logout", web_controllers.Logout()).Name("logout")
	authRouter.Get("profile", web_controllers.Profile()).Name("profile")
	authRouter.Get("posts", web_controllers.PostIndex(db)).Name("posts.index")

	fmt.Println("web route register success")
}
