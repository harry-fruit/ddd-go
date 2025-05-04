package httphandler

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/harry-fruit/ddd-go/internal/domain/service"
	httpserver "github.com/harry-fruit/ddd-go/internal/infrastructure/api/http"
)

type ProductHandler struct {
	service       *service.ProductService
	routeHandlers []*httpserver.RouteHandler
}

func (ph *ProductHandler) GetRoutes() []*httpserver.RouteHandler {
	return ph.routeHandlers
}

func (ph *ProductHandler) CreateProduct(c *fiber.Ctx) error {
	fmt.Println(c.Method())

	return c.SendString("Hello, World!")
}

func (ph *ProductHandler) GetProducts(c *fiber.Ctx) error {
	fmt.Println(c.Method())

	return c.SendString("Hello, World!")
}

func (ph *ProductHandler) GetProductById(c *fiber.Ctx) error {
	fmt.Println("DI Worked just fine")
	return nil
}

func (ph *ProductHandler) UpdateProduct(c *fiber.Ctx) error {
	fmt.Println(c.Method())
	return c.SendString("Hello, World!")
}

func (ph *ProductHandler) DeleteProduct(c *fiber.Ctx) error {
	fmt.Println(c.Method())
	return c.SendString("Hello, World!")
}

func NewProductHandler(service *service.ProductService) httpserver.Handler {
	productHandler := &ProductHandler{
		service: service,
	}

	fmt.Println("NewProductHandler")

	productHandler.routeHandlers = []*httpserver.RouteHandler{
		httpserver.NewRouteHandler(http.MethodGet, "/products", productHandler.GetProducts),
		httpserver.NewRouteHandler(http.MethodPost, "/products", productHandler.CreateProduct),
		httpserver.NewRouteHandler(http.MethodGet, "/products/:id", productHandler.GetProductById),
		httpserver.NewRouteHandler(http.MethodPatch, "/products/:id", productHandler.UpdateProduct),
		httpserver.NewRouteHandler(http.MethodDelete, "/products/:id", productHandler.DeleteProduct),
	}

	return productHandler
}
