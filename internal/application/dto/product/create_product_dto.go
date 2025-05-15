package productdto

import (
	"github.com/harry-fruit/ddd-go/internal/domain/entity"
)

type CreateProductDTO struct {
	UniqueKey      string   `json:"unique_key" validate:"required,min=3,max=150"`
	Name           string   `json:"name" validate:"required,min=3,max=150"`
	Description    *string  `json:"description,omitempty"`
	Price          float64  `json:"price" validate:"required"`
	RemainingStock *int     `json:"remaining_stock" validate:"required"`
	Rating         *float64 `json:"rating,omitempty"`
}

func (dto *CreateProductDTO) ToDomain() *entity.Product {
	return entity.NewProduct(dto.UniqueKey, dto.Name, dto.Description, dto.Price, dto.RemainingStock, dto.Rating)
}
