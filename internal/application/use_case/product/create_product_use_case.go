package productusecase

import (
	"github.com/harry-fruit/ddd-go/internal/domain/entity"
	abstractrepository "github.com/harry-fruit/ddd-go/internal/domain/repository"
)

type Params = entity.Product
type Result = bool

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

func (uc *CreateProductUseCase) Execute(params *Params) (result Result, err error) {
	error := uc.productRepository.CreateProduct(params)

	if error != nil {
		return false, error
	}

	return true, nil
}
