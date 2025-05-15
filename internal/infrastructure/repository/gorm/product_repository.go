package gormrepository

import (
	"errors"

	productdto "github.com/harry-fruit/ddd-go/internal/application/dto/product"
	"github.com/harry-fruit/ddd-go/internal/domain/entity"
	abstractrepository "github.com/harry-fruit/ddd-go/internal/domain/repository"
	gormmapper "github.com/harry-fruit/ddd-go/internal/infrastructure/mapper/gorm"
	gormmodel "github.com/harry-fruit/ddd-go/internal/infrastructure/model/gorm"
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

func (pr *ProductRepository) Create(product *entity.Product) (uint, error) {
	payload := gormmapper.NewProductMapper().ToDB(*product)
	err := pr.db.Create(&payload).Error

	if err != nil {
		return 0, err
	}

	return payload.ID, nil
}

func (pr *ProductRepository) Update(product productdto.UpdateProductDTO) (int, error) {
	return 0, nil
}

func (pr *ProductRepository) Delete(id int) error {
	return nil
}

func (pr *ProductRepository) Exists(uniqueKey string) (bool, error) {
	entity := &gormmodel.ProductModel{}
	err := pr.db.First(entity, "unique_key = ?", uniqueKey).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
