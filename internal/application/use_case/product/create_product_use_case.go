package productusecase

import (
	"fmt"

	"github.com/harry-fruit/ddd-go/internal/domain/entity"
	abstractrepository "github.com/harry-fruit/ddd-go/internal/domain/repository"
	apperror "github.com/harry-fruit/ddd-go/internal/infrastructure/errors"
)

type CreateProductUseCaseParamsResult = uint

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

func (uc *CreateProductUseCase) Execute(params *entity.Product) (result CreateProductUseCaseParamsResult, err error) {
	exists, err := uc.productRepository.Exists(params.UniqueKey)

	if err != nil {
		return 0, apperror.NewInternalServerError(err)
	}

	if exists {
		return 0, apperror.NewBadRequestError(fmt.Sprintf("product with unique key '%s' already exists", params.UniqueKey))
	}

	id, error := uc.productRepository.Create(params)

	if error != nil {
		return 0, apperror.NewInternalServerError(error)
	}

	return id, nil
}
