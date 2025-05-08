package di

import (
	"log"

	"github.com/harry-fruit/ddd-go/config"
	productusecase "github.com/harry-fruit/ddd-go/internal/application/use_case/product"
	"github.com/harry-fruit/ddd-go/internal/domain/service"
	httpserver "github.com/harry-fruit/ddd-go/internal/infrastructure/api/http"
	httphandler "github.com/harry-fruit/ddd-go/internal/infrastructure/api/http/handler"
	"github.com/harry-fruit/ddd-go/internal/infrastructure/api/http/router"
	infrastructure "github.com/harry-fruit/ddd-go/internal/infrastructure/database"
	"github.com/harry-fruit/ddd-go/internal/infrastructure/repository"
	"go.uber.org/dig"
)

func BuildContainer() *dig.Container {
	container := dig.New()

	container.Provide(func() *config.Config {
		cfg, err := config.LoadConfig()
		if err != nil {
			log.Fatalf("failed to load config %v", err)
		}
		return cfg
	})

	// Relational DB
	container.Provide(infrastructure.NewPostgresDB)

	// Repositories
	container.Provide(repository.NewProductRepository)

	// Use Cases
	container.Provide(productusecase.NewGetProductsUseCase)
	container.Provide(productusecase.NewGetProductByIdProductUseCase)
	container.Provide(productusecase.NewCreateProductUseCase)
	container.Provide(productusecase.NewUpdateProductUseCase)
	container.Provide(productusecase.NewDeleteProductUseCase)

	// Services
	container.Provide(service.NewProductService)

	// Handlers
	// container.Provide(httphandler.NewProductHandler, dig.Group("handlers"))
	container.Provide(httphandler.NewProductHandler)

	// Routes
	container.Provide(router.NewProductRoute, dig.Group("routes"))

	// API
	container.Provide(httpserver.NewHTTPServer)

	return container
}
