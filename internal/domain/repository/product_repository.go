package abstractrepository

import "github.com/harry-fruit/ddd-go/internal/domain/entity"

type ProductRepository interface {
	GetProducts() ([]*entity.Product, error)
	GetProductById(id int) (*entity.Product, error)
	CreateProduct(product *entity.Product) error
	UpdateProduct(product *entity.Product) error
	DeleteProduct(product *entity.Product) error
}
