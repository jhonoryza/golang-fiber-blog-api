package main

import (
	"fiber_blog/config"
	"fiber_blog/providers/route"
	"fiber_blog/routes"
)

func main() {
	config.InitLogging()
	db := config.InitDatabase()

	router := routes.Initialize()

	routes.RegisterApiRoute(router, db)
	routes.RegisterWebRoute(router, db)

	route.LoadRouteMap(router)

	routes.AppListen(router)
}
