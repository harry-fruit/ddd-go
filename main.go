package main

import (
	"fmt"
	"log"

	// "net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/harry-fruit/ddd-go/config"
	productusecase "github.com/harry-fruit/ddd-go/internal/application/use_case/product"
	"github.com/harry-fruit/ddd-go/internal/domain/service"

	// httpserver "github.com/harry-fruit/ddd-go/internal/infrastructure/api/http"
	// httphandler "github.com/harry-fruit/ddd-go/internal/infrastructure/api/http/handler"
	// "github.com/harry-fruit/ddd-go/internal/infrastructure/api/http/router"
	httphandler "github.com/harry-fruit/ddd-go/internal/infrastructure/api/http/handler"
	"github.com/harry-fruit/ddd-go/internal/infrastructure/api/http/router"
	pgdb "github.com/harry-fruit/ddd-go/internal/infrastructure/database"
	"github.com/harry-fruit/ddd-go/internal/infrastructure/repository"
)

func main() {
	fmt.Println("Starting server...")

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config %v", err)
	}

	app := fiber.New(fiber.Config{
		AppName: "ddd-go",
	})

	api := app.Group("/api")
	v1 := api.Group("/v1")

	// Create database
	db, err := pgdb.NewPostgresDB(cfg)
	if err != nil {
		log.Fatalf("failed to create db %v", err)
	}

	// Create Repository
	productRepository := repository.NewProductRepository(db)

	// Create usecases
	createProductUseCase := productusecase.NewCreateProductUseCase(productRepository)
	getProductByIdUseCase := productusecase.NewGetProductByIdProductUseCase(productRepository)
	getProductsUseCase := productusecase.NewGetProductsUseCase(productRepository)
	updateProductUseCase := productusecase.NewUpdateProductUseCase(productRepository)
	deleteProductUseCase := productusecase.NewDeleteProductUseCase(productRepository)

	// Create services
	productService := service.NewProductService(
		productRepository,
		createProductUseCase,
		getProductByIdUseCase,
		getProductsUseCase,
		updateProductUseCase,
		deleteProductUseCase,
	)

	// Create handlers
	productHandler := httphandler.NewProductHandler(productService)

	// Register routes
	router.RegisterProductRoutes(v1, productHandler)
	fmt.Println("Server running at:", cfg.ServerPort)

	port := fmt.Sprintf(":%s", cfg.ServerPort)
	e := app.Listen(port)

	if e != nil {
		log.Fatalf("Fiber failed to start: %v", e)
	}
}
