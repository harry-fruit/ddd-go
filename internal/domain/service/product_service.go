package service

import (
	productdto "github.com/harry-fruit/ddd-go/internal/application/dto/product"
	productusecase "github.com/harry-fruit/ddd-go/internal/application/use_case/product"
	abstractrepository "github.com/harry-fruit/ddd-go/internal/domain/repository"
	pagination "github.com/harry-fruit/ddd-go/pkg/pagination"

	"github.com/harry-fruit/ddd-go/internal/domain/entity"
)

type ProductService struct {
	productRepository     abstractrepository.ProductRepository
	createProductUseCase  *productusecase.CreateProductUseCase
	getProductByIdUseCase *productusecase.GetProductByIdProductUseCase
	getProductsUseCase    *productusecase.GetProductsUseCase
	updateProductUseCase  *productusecase.UpdateProductUseCase
	deleteProductUseCase  *productusecase.DeleteProductUseCase
}

func NewProductService(
	productRepository abstractrepository.ProductRepository,
	createProductUseCase *productusecase.CreateProductUseCase,
	getProductByIdUseCase *productusecase.GetProductByIdProductUseCase,
	getProductsUseCase *productusecase.GetProductsUseCase,
	updateProductUseCase *productusecase.UpdateProductUseCase,
	deleteProductUseCase *productusecase.DeleteProductUseCase,
) *ProductService {
	return &ProductService{
		productRepository:     productRepository,
		createProductUseCase:  createProductUseCase,
		getProductByIdUseCase: getProductByIdUseCase,
		getProductsUseCase:    getProductsUseCase,
		updateProductUseCase:  updateProductUseCase,
		deleteProductUseCase:  deleteProductUseCase,
	}
}

func (ps *ProductService) CreateProduct(params *productdto.CreateProductDTO) (uint, error) {
	return ps.createProductUseCase.Execute(params.ToDomain())
}

func (ps *ProductService) GetProductById(id int) (*entity.Product, error) {
	return ps.getProductByIdUseCase.Execute(id)
}

func (ps *ProductService) GetProducts(params *pagination.PaginationParams) (*pagination.PaginationResult[entity.Product], error) {
	return ps.getProductsUseCase.Execute(params)
}

func (ps *ProductService) UpdateProduct(params productdto.UpdateProductDTO) (int, error) {
	return ps.updateProductUseCase.Execute(params)
}

func (ps *ProductService) DeleteProduct(id int) (bool, error) {
	return ps.deleteProductUseCase.Execute(id)
}
