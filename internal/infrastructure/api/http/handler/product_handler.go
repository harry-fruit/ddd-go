package httphandler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	productdto "github.com/harry-fruit/ddd-go/internal/application/dto/product"
	"github.com/harry-fruit/ddd-go/internal/domain/service"
	httpserver "github.com/harry-fruit/ddd-go/internal/infrastructure/api/http"
	apperror "github.com/harry-fruit/ddd-go/internal/infrastructure/errors"
	sharedinfra "github.com/harry-fruit/ddd-go/internal/shared/infrastructure"
	"github.com/harry-fruit/ddd-go/pkg/pagination"
)

type ProductHandler struct {
	service *service.ProductService
}

func (ph *ProductHandler) CreateProduct(c *fiber.Ctx) error {
	var req productdto.CreateProductDTO

	if err := c.BodyParser(&req); err != nil {
		return httpserver.NewHTTPResponseBadRequest(c, err.Error())
	}

	validate := sharedinfra.NewValidate()

	if errs := validate.Execute(&req); errs != nil {
		return httpserver.NewHTTPResponseBadRequest(c, errs)
	}

	id, err := ph.service.CreateProduct(&req)

	if err != nil {
		return apperror.Handle(c, err)
	}

	response := map[string]any{
		"id": id,
	}

	return httpserver.NewHTTPResponseCreated(c, response)
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
