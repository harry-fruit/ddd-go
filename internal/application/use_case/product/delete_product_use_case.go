package productusecase

import (
	abstractrepository "github.com/harry-fruit/ddd-go/internal/domain/repository"
)

type DeleteProductUseCaseParams = int
type DeleteProductUseCaseParamsResult = bool

type DeleteProductUseCase struct {
	productRepository abstractrepository.ProductRepository
}

func NewDeleteProductUseCase(
	productRepository abstractrepository.ProductRepository,
) *DeleteProductUseCase {
	return &DeleteProductUseCase{
		productRepository: productRepository,
	}
}

func (uc *DeleteProductUseCase) Execute(params DeleteProductUseCaseParams) (result DeleteProductUseCaseParamsResult, err error) {
	error := uc.productRepository.Delete(params)

	if error != nil {
		return false, error
	}

	return true, nil
}
