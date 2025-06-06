package abstractrepository

import (
	productdto "github.com/harry-fruit/ddd-go/internal/application/dto/product"
	pagination "github.com/harry-fruit/ddd-go/pkg/pagination"

	"github.com/harry-fruit/ddd-go/internal/domain/entity"
)

type ProductRepository interface {
	Get(params *pagination.PaginationParams) (*pagination.PaginationResult[entity.Product], error)
	GetById(id int) (*entity.Product, error)
	Create(product *entity.Product) (id uint, err error)
	Update(product productdto.UpdateProductDTO) (id int, err error)
	Exists(uniqueKey string) (bool, error)
	Delete(id int) error
}
