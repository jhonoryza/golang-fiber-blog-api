package route

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

var routeMap = make(map[string]string)

func LoadRouteMap(router *fiber.App) {
	allRoutes := router.GetRoutes()
	for _, route := range allRoutes {
		if route.Method == "HEAD" || route.Name == "" {
			continue
		}
		routeMap[route.Name] = route.Path
		fmt.Printf("%s -> %s\n", route.Name, route.Path)
	}
}

func GetRouteURL(routeName string) string {
	url, exists := routeMap[routeName]
	if !exists {
		_ = fmt.Errorf("route %s not found", routeName)
	}
	return url
}
