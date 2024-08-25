package main

import (
	"fiber_blog/config"
	"fiber_blog/routes"
)

func main() {
	config.InitLogging()
	db := config.InitDatabase()

	router := routes.Initialize()

	routes.RegisterApiRoute(router, db)
	routes.RegisterWebRoute(router, db)

	routes.AppListen(router)
}
