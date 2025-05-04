package router

import (
	"github.com/gofiber/fiber/v2"
	httphandler "github.com/harry-fruit/ddd-go/internal/infrastructure/api/http/handler"
)

func RegisterProductRoutes(app fiber.Router, productHandler *httphandler.ProductHandler) {
	app.Get("/products", productHandler.GetProducts)
	app.Post("/products", productHandler.CreateProduct)
	app.Get("/products/:id", productHandler.GetProductById)
	app.Put("/products/:id", productHandler.UpdateProduct)
	app.Delete("/products/:id", productHandler.DeleteProduct)
}
