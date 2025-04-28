package productdto

type CreateProductDTO struct {
	Id             int
	UniqueKey      *string  `json:"unique_key"`
	Name           *string  `json:"name"`
	Description    *string  `json:"description,omitempty"`
	Price          *float64 `json:"price"`
	RemainingStock *int     `json:"remaining_stock"`
	Rating         *float64 `json:"rating"`
}
