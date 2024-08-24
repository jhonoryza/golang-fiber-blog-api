package config

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	fiberRecover "github.com/gofiber/fiber/v2/middleware/recover"
	"os"
	"os/signal"
)

var router *fiber.App

func InitRouter() {
	inertia := GetInertia()
	router = fiber.New(fiber.Config{
		Views: inertia,
	})

	router.Use(inertia.Middleware())

	router.Use(favicon.New(favicon.Config{
		File: "./public/favicon.ico",
	}))

	router.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))

	router.Use(fiberRecover.New())

	fmt.Println("router initialized")
}

func GetRouter() *fiber.App {
	if router == nil {
		InitRouter()
	}
	return router
}

func AppListen() {
	router := GetRouter()
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
	log.Info("Listening on port " + GetEnv().GetString("APP_PORT"))
	fmt.Println("Listening on port " + GetEnv().GetString("APP_PORT"))
	err := router.Listen(":" + GetEnv().GetString("APP_PORT"))
	if err != nil {
		log.Error(err)
		fmt.Println(err)
		err = router.Shutdown()
		if err != nil {
			panic(err)
		}
	}
}
