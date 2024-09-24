package routes

import (
	"errors"
	"fiber_blog/env"
	"fmt"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/jhonoryza/inertia-fiber"
)

func fixMimeType(router *fiber.App) {
	router.All("/build/*.js", func(ctx *fiber.Ctx) error {
		ctx.Set(fiber.HeaderContentType, "text/javascript")
		return ctx.Next()
	})
	router.All("/build/*.css", func(ctx *fiber.Ctx) error {
		ctx.Set(fiber.HeaderContentType, "text/css")
		return ctx.Next()
	})
}

func initRouter() *fiber.App {
	return fiber.New(fiber.Config{
		// Override default error handler
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			isJsonRequest := ctx.Get("Accept") == "application/json" || ctx.Get("Content-Type") == "application/json"

			// Status code defaults to 500
			code := fiber.StatusInternalServerError

			// Retrieve the custom status code if it's a *fiber.Error
			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}
			if code == fiber.StatusNotFound {
				fmt.Println(e.Error())
				log.Error(e.Error())
				if isJsonRequest {
					return ctx.JSON(fiber.Map{
						"code":    code,
						"message": e.Error(),
						"data":    nil,
					})
				}
				return inertia.Render(ctx, 404, "Error", fiber.Map{
					"message": "Page Not Found",
				})
			}

			fmt.Println(e.Error())
			log.Error(e.Error())
			if isJsonRequest {
				return ctx.JSON(fiber.Map{
					"code":    code,
					"message": e.Error(),
					"data":    nil,
				})
			}
			return inertia.Render(ctx, 500, "Error", fiber.Map{
				"message": "Internal Server Error",
			})
		},
	})
}

func Initialize() *fiber.App {
	router := initRouter()

	router.Use(favicon.New(favicon.Config{
		File: "./public/favicon.ico",
	}))

	router.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))

	router.Use(csrf.New(csrf.Config{
		SingleUseToken: true,
	}))

	//router.Use(helmet.New(helmet.Config{
	//	ReferrerPolicy: "strict-origin-when-cross-origin",
	//}))
	r := inertia.NewRenderer("app")

	r.MustParseGlob("resources/views/*.html")
	r.ViteBasePath = "/build/"
	r.AddViteEntryPoint("resources/js/app.js")
	r.MustParseViteManifestFile("public/build/manifest.json")

	router.Use(inertia.Middleware(r))

	router.Use(logger.New())
	// router.Use(recover.New())

	fixMimeType(router)

	//router.Static("/assets", "public/build/assets")
	router.Static("/build/assets", "public/build/assets")
	router.Static("/storage", "public/storage")

	fmt.Println("router initialized")
	return router
}

func AppListen(router *fiber.App) {
	// Close any connections on interrupt signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		log.Info("app shutting down")
		err := router.Shutdown()
		if err != nil {
			log.Error(err.Error())
			fmt.Println(err)
			panic(err)
		}
	}()

	// Start listening on the specified address
	log.Info("Listening on port " + env.GetEnv().GetString("APP_PORT"))
	fmt.Println("Listening on port " + env.GetEnv().GetString("APP_PORT"))
	err := router.Listen(":" + env.GetEnv().GetString("APP_PORT"))
	if err != nil {
		log.Error(err)
		fmt.Println(err)
		err = router.Shutdown()
		if err != nil {
			panic(err)
		}
	}
}
