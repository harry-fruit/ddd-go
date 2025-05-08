package sharedinfra

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/dig"
)

type RoutesIn struct {
	dig.In
	Routes []Route `group:"routes"`
}

type Route interface {
	Register(fiber.Router)
}
