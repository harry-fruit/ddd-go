package productusecase

import (
	"github.com/harry-fruit/ddd-go/internal/domain/entity"
	abstractrepository "github.com/harry-fruit/ddd-go/internal/domain/repository"
)

type GetProductByIdUseCaseParams = int
type GetProductByIdseCaseResult = *entity.Product

type GetProductByIdProductUseCase struct {
	productRepository abstractrepository.ProductRepository
}

func NewGetProductByIdProductUseCase(
	productRepository abstractrepository.ProductRepository,
) *GetProductByIdProductUseCase {
	return &GetProductByIdProductUseCase{
		productRepository: productRepository,
	}
}

func (uc *GetProductByIdProductUseCase) Execute(params GetProductByIdUseCaseParams) (result GetProductByIdseCaseResult, err error) {
	product, error := uc.productRepository.GetById(params)

	if error != nil {
		return nil, error
	}

	return product, nil
}
