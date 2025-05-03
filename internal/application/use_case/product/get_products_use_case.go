package productusecase

import (
	abstractrepository "github.com/harry-fruit/ddd-go/internal/domain/repository"
	pagination "github.com/harry-fruit/ddd-go/pkg/pagination"

	"github.com/harry-fruit/ddd-go/internal/domain/entity"
)

type GetProductUseCaseParams = pagination.PaginationParams
type GetProductUseCaseResult = *pagination.PaginationResult[[]*entity.Product]

type GetProductsUseCase struct {
	productRepository abstractrepository.ProductRepository
}

func NewGetProductsUseCase(
	productRepository abstractrepository.ProductRepository,
) *GetProductsUseCase {
	return &GetProductsUseCase{
		productRepository: productRepository,
	}
}

func (uc *GetProductsUseCase) Execute(params *GetProductUseCaseParams) (result GetProductUseCaseResult, err error) {
	products, error := uc.productRepository.Get(params)

	if error != nil {
		return nil, error
	}

	return products, nil
}
