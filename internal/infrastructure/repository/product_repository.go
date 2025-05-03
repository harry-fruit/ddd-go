package repository

import (
	"database/sql"

	productdto "github.com/harry-fruit/ddd-go/internal/application/dto/product"
	"github.com/harry-fruit/ddd-go/internal/domain/entity"
	"github.com/harry-fruit/ddd-go/pkg/pagination"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (pr *ProductRepository) Get(params *pagination.PaginationParams) (*pagination.PaginationResult[[]*entity.Product], error) {
	return nil, nil
}

func (pr *ProductRepository) GetById(id int) (*entity.Product, error) {
	query := "SELECT id, unique_key, name, description, price, remaining_stock FROM product WHERE id = $1"

	row := pr.db.QueryRow(query, id)

	var entity entity.Product
	if err := row.Scan(&entity.Id, &entity.UniqueKey, &entity.Name, &entity.Description, &entity.Price, &entity.RemainingStock); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &entity, nil
}

func (pr *ProductRepository) Create(product productdto.CreateProductDTO) (int, error) {
	query := `
		INSERT INTO product (unique_key, name, description, price, remaining_stock)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`

	var id int
	err := pr.db.QueryRow(query, product.UniqueKey, product.Name, product.Description, product.Price, product.RemainingStock).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (pr *ProductRepository) Update(product productdto.UpdateProductDTO) (int, error) {
	return 0, nil
}

func (pr *ProductRepository) Delete(id int) error {
	return nil
}
