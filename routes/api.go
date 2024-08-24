package routes

import (
	"fiber_blog/app/controllers"
	"fiber_blog/config"
	"fmt"
)

func RegisterApiRoute() {
	router := config.GetRouter()
	api := router.Group("/api")

	api.Get("/posts", controllers.PostsIndex)
	api.Get("/posts/:slug", controllers.PostsShow)

	fmt.Println("api route register success")
}
