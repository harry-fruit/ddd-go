package gormrepository

import (
	productdto "github.com/harry-fruit/ddd-go/internal/application/dto/product"
	"github.com/harry-fruit/ddd-go/internal/domain/entity"
	abstractrepository "github.com/harry-fruit/ddd-go/internal/domain/repository"
	"github.com/harry-fruit/ddd-go/pkg/pagination"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) abstractrepository.ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (pr *ProductRepository) Get(params *pagination.PaginationParams) (*pagination.PaginationResult[entity.Product], error) {
	//TODO:
	return &pagination.PaginationResult[entity.Product]{}, nil
}

func (pr *ProductRepository) GetById(id int) (*entity.Product, error) {
	//TODO:
	return &entity.Product{}, nil
}

func (pr *ProductRepository) Create(product productdto.CreateProductDTO) (int, error) {
	return 0, nil
}

func (pr *ProductRepository) Update(product productdto.UpdateProductDTO) (int, error) {
	return 0, nil
}

func (pr *ProductRepository) Delete(id int) error {
	return nil
}
