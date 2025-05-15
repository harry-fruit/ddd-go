package httpserver

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/harry-fruit/ddd-go/config"
	sharedinfra "github.com/harry-fruit/ddd-go/internal/shared/infrastructure"

	_ "github.com/harry-fruit/ddd-go/docs"
)

type HTTPServer struct {
	port   string
	Routes []sharedinfra.Route
}

func NewHTTPServer(routes sharedinfra.RoutesIn, config *config.Config) *HTTPServer {
	return &HTTPServer{
		port:   fmt.Sprintf(":%s", config.ServerPort),
		Routes: routes.Routes,
	}
}

// @title DDD-Go
// @version 0.0.1
// @description This is a sample Rest API for Golang that uses DDD, Clean Architecture, and SOLID principles.
// @contact.name Isaque
// @contact.email contact@isaquedev.com
// @license.name Apache 2.0 // TODO: Change this
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html // TODO: Change this
// @host localhost:3001 // TODO: Change this
// @BasePath /api/v1 // TODO: Change this
func (s *HTTPServer) Start() {

	app := fiber.New(fiber.Config{
		AppName: "ddd-go",
	})

	api := app.Group("/api")
	v1 := api.Group("/v1")

	app.Get("/swagger/*", swagger.New(swagger.Config{
		URL:          "/swagger/doc.json", // or where swag outputs your spec
		DocExpansion: "none",
	}))

	for _, route := range s.Routes {
		route.Register(v1)
	}

	e := app.Listen(s.port)

	if e != nil {
		log.Fatalf("Fiber failed to start: %v", e)
	}
}
