package abstractrepository

import (
	productdto "github.com/harry-fruit/ddd-go/internal/application/dto/product"
	pagination "github.com/harry-fruit/ddd-go/pkg"

	"github.com/harry-fruit/ddd-go/internal/domain/entity"
)

type ProductRepository interface {
	Get(params *pagination.PaginationParams) (*pagination.PaginationResult[[]*entity.Product], error)
	GetById(id int) (*entity.Product, error)
	Create(product productdto.CreateProductDTO) error
	Update(product productdto.UpdateProductDTO) error
	Delete(id int) error
}
