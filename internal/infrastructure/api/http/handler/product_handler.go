package httphandler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/harry-fruit/ddd-go/internal/domain/service"
)

type ProductHandler struct {
	service *service.ProductService
}

func NewProductHandler(service *service.ProductService) *ProductHandler {
	productHandler := &ProductHandler{
		service: service,
	}
	return productHandler
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
	fmt.Println(c.Method())
	return c.SendString("Hello, World!")
}

func (ph *ProductHandler) UpdateProduct(c *fiber.Ctx) error {
	fmt.Println(c.Method())
	return c.SendString("Hello, World!")
}

func (ph *ProductHandler) DeleteProduct(c *fiber.Ctx) error {
	fmt.Println(c.Method())
	return c.SendString("Hello, World!")
}
