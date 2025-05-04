package httpserver

import "github.com/gofiber/fiber/v2"

type RouteHandler struct {
	Method  string
	Path    string
	Handler func(c *fiber.Ctx) error
}

func NewRouteHandler(method string, path string, handler func(c *fiber.Ctx) error) *RouteHandler {
	return &RouteHandler{
		Method:  method,
		Path:    path,
		Handler: handler,
	}
}
