package productusecase

import (
	productdto "github.com/harry-fruit/ddd-go/internal/application/dto/product"
	abstractrepository "github.com/harry-fruit/ddd-go/internal/domain/repository"
)

type UpdateProductUseCaseParamsResult = int

type UpdateProductUseCase struct {
	productRepository abstractrepository.ProductRepository
}

func NewUpdateProductUseCase(
	productRepository abstractrepository.ProductRepository,
) *UpdateProductUseCase {
	return &UpdateProductUseCase{
		productRepository: productRepository,
	}
}

func (uc *UpdateProductUseCase) Execute(params productdto.UpdateProductDTO) (result UpdateProductUseCaseParamsResult, err error) {
	id, error := uc.productRepository.Update(params)

	if error != nil {
		return id, error
	}

	return id, nil
}
