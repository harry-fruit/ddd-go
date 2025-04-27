package service

import (
	productusecase "github.com/harry-fruit/ddd-go/internal/application/use_case/product"
	"github.com/harry-fruit/ddd-go/internal/domain/entity"
	abstractrepository "github.com/harry-fruit/ddd-go/internal/domain/repository"
)

type ProductService struct {
	productRepository    abstractrepository.ProductRepository
	createProductUseCase *productusecase.CreateProductUseCase
}

func NewProductService(
	productRepository abstractrepository.ProductRepository,
	createProductUseCase *productusecase.CreateProductUseCase,
) *ProductService {
	return &ProductService{
		productRepository:    productRepository,
		createProductUseCase: createProductUseCase,
	}
}

func (ps *ProductService) CreateProduct(product *entity.Product) (bool, error) {
	return ps.createProductUseCase.Execute(product)
}
