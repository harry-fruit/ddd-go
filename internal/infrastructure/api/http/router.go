package httpserver

import (
	"github.com/gofiber/fiber/v2"
)

// TODO: Try to reduce this to reach N(1) instead of N(O)
func RegisterAllRoutes(app fiber.Router, handlers []Handler) {
	for _, handler := range handlers {
		for _, routeHandler := range handler.GetRoutes() {
			app.Add(routeHandler.Method, routeHandler.Path, routeHandler.Handler)
		}
	}
}
