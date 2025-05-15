package router

import (
	"github.com/gofiber/fiber/v2"
	httphandler "github.com/harry-fruit/ddd-go/internal/infrastructure/api/http/handler"
	sharedinfra "github.com/harry-fruit/ddd-go/internal/shared/infrastructure"
)

type ProductRoute struct {
	ProductHandler *httphandler.ProductHandler
}

func NewProductRoute(productHandler *httphandler.ProductHandler) sharedinfra.Route {
	return &ProductRoute{
		ProductHandler: productHandler,
	}
}

func (pr *ProductRoute) Register(app fiber.Router) {
	app.Get("/products", pr.ProductHandler.GetProducts)
	app.Post("/products", pr.ProductHandler.CreateProduct)
	app.Get("/products/:id", pr.ProductHandler.GetProductById)
	app.Put("/products/:id", pr.ProductHandler.UpdateProduct)
	app.Delete("/products/:id", pr.ProductHandler.DeleteProduct)
}
