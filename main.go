package main

import (
	_ "embed"
	"fiber_blog/app"
	"fiber_blog/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
)

func main() {
	// setup log file
	appLog, err := os.OpenFile("app.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	// setup logger
	log.SetOutput(appLog)

	// setup config
	config := app.GetConfig()

	// setup fiber
	router := fiber.New()

	db, err := gorm.Open(postgres.Open(config.GetString("DATABASE_URL")), &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Error(err)
		panic(err)
	}

	// setup router
	router.Get("/", controllers.HomeIndex)
	postController := controllers.NewPostController(db)
	router.Get("/api/posts", postController.Index)
	router.Get("/api/posts/:slug", postController.Show)

	// setup server listen
	log.Info("Listening on port " + config.GetString("APP_PORT"))
	err = router.Listen(":" + config.GetString("APP_PORT"))
	if err != nil {
		log.Error(err)
		panic(err)
	}
}
