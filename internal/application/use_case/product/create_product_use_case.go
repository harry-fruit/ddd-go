package productusecase

import (
	productdto "github.com/harry-fruit/ddd-go/internal/application/dto/product"
	abstractrepository "github.com/harry-fruit/ddd-go/internal/domain/repository"
)

type CreateProductUseCaseParamsResult = int

type CreateProductUseCase struct {
	productRepository abstractrepository.ProductRepository
}

func NewCreateProductUseCase(
	productRepository abstractrepository.ProductRepository,
) *CreateProductUseCase {
	return &CreateProductUseCase{
		productRepository: productRepository,
	}
}

func (uc *CreateProductUseCase) Execute(params productdto.CreateProductDTO) (result CreateProductUseCaseParamsResult, err error) {
	id, error := uc.productRepository.Create(params)

	if error != nil {
		return id, error
	}

	return id, nil
}
