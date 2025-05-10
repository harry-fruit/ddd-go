package httphandler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/harry-fruit/ddd-go/internal/domain/service"
	"github.com/harry-fruit/ddd-go/pkg/pagination"
)

type ProductHandler struct {
	service *service.ProductService
}

func (ph *ProductHandler) CreateProduct(c *fiber.Ctx) error {
	fmt.Println(c.Method())

	return c.SendString("Hello, World!")
}

// GetProducts godoc
// @Summary      List all products
// @Description  Get a list of all products in the catalog
// @Tags         products
// @Accept       json
// @Produce      json
// @Param    	 page        query   int     false  "Page number"
// @Param        page_size   query   int     false  "Items per page"
// @Success      200  {object}  pagination.PaginationResult[entity.Product]
// @Failure      500  {object}  nil
// @Router       /products [get]
func (ph *ProductHandler) GetProducts(c *fiber.Ctx) error {
	var params pagination.PaginationParams
	if err := c.QueryParser(&params); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	result, err := ph.service.GetProducts(&params)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(result)
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

func NewProductHandler(service *service.ProductService) *ProductHandler {
	productHandler := &ProductHandler{
		service: service,
	}

	return productHandler
}
