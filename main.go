package main

import (
	"fiber_blog/config"
	"fiber_blog/routes"
)

func main() {
	config.InitLogging()
	config.InitDatabase()
	config.InitInertia()
	config.InitRouter()

	routes.RegisterApiRoute()
	routes.RegisterWebRoute()

	config.AppListen()
}
