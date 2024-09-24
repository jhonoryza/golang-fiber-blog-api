package routes

import (
	"encoding/json"
	"fiber_blog/app/controllers/web_controllers"
	"fiber_blog/env"
	"fiber_blog/utils"
	"fmt"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jhonoryza/inertia-fiber"
	"golang.org/x/exp/maps"
	"gorm.io/gorm"
)

func RegisterWebRoute(router *fiber.App, db *gorm.DB) {
	store := session.New()

	router.Use(func(ctx *fiber.Ctx) error {
		sess, err := store.Get(ctx)
		if err != nil {
			panic(err)
		}
		var flashMessage utils.FlashMessage
		flash := sess.Get("flash")
		if flash != nil {
			_ = json.Unmarshal(flash.([]byte), &flashMessage)
			sess.Delete("flash")
			_ = sess.Save()
		}
		//fmt.Printf("message: %v, type: %v\n", flashMessage.Message, flashMessage.Type)
		flashProps := fiber.Map{
			"flash": utils.FlashMessage{
				Message: flashMessage.Message,
				Type:    flashMessage.Type,
			},
		}
		inertia.Share(ctx, flashProps)
		return ctx.Next()
	})

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

				parentProps := inertia.Shared(ctx)
				userProps := fiber.Map{
					"name": name,
				}
				maps.Copy(userProps, parentProps)

				inertia.Share(ctx, userProps)

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
	router.Post("/login", web_controllers.Login(db, store)).Name("login")

	authRouter := router.Group("/auth").Name("auth.")

	// redirect to log in form if cookie is empty
	authRouter.Use(func(c *fiber.Ctx) error {
		cookie := c.Cookies(env.GetEnv().GetString("COOKIE_NAME"))
		if cookie == "" {
			return inertia.RedirectToRoute(c, "login.form", fiber.Map{})
		}

		// save token in context tobe used in the nex middleware
		c.Locals("user", cookie)
		return c.Next()
	})

	// redirect to log in form if not authenticated
	authRouter.Use(jwtware.New(jwtware.Config{
		SigningKey:  jwtware.SigningKey{Key: []byte(env.GetEnv().GetString("JWT_SECRET"))},
		TokenLookup: "cookie:" + env.GetEnv().GetString("COOKIE_NAME"),
		ContextKey:  "user",
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return inertia.RedirectToRoute(c, "login.form", fiber.Map{})
		},
	}))

	authRouter.Get("dashboard", web_controllers.Dashboard()).Name("dashboard")
	authRouter.Post("logout", web_controllers.Logout(store)).Name("logout")
	authRouter.Get("profile", web_controllers.Profile()).Name("profile")
	authRouter.Get("posts", web_controllers.PostIndex(db)).Name("posts.index")

	fmt.Println("web route register success")
}
