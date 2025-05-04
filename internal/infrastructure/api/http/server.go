package httpserver

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/harry-fruit/ddd-go/config"
	"go.uber.org/dig"
)

type HTTPServer struct {
	port     string
	Handlers []Handler
}

type HandlersIn struct {
	dig.In
	Handlers []Handler `group:"handlers"`
}

func NewHTTPServer(h HandlersIn, config *config.Config) *HTTPServer {
	return &HTTPServer{
		port:     fmt.Sprintf(":%s", config.ServerPort),
		Handlers: h.Handlers,
	}
}

func (s *HTTPServer) Start(container *dig.Container) {
	app := fiber.New(fiber.Config{
		AppName: "ddd-go",
	})

	container.Invoke(RegisterAllRoutes)

	api := app.Group("/api")
	v1 := api.Group("/v1")

	RegisterAllRoutes(v1, s.Handlers)

	e := app.Listen(s.port)

	if e != nil {
		log.Fatalf("Fiber failed to start: %v", e)
	}
}
