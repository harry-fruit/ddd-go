package productdto

type CreateProductDTO struct {
	UniqueKey      string   `json:"unique_key" validate:"required,min=3,max=150"`
	Name           string   `json:"name" validate:"required,min=3,max=150"`
	Description    *string  `json:"description,omitempty"`
	Price          float64  `json:"price" validate:"required"`
	RemainingStock int      `json:"remaining_stock" validate:"required"`
	Rating         *float64 `json:"rating,omitempty"`
}
