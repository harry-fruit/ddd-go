package productdto

type UpdateProductDTO struct {
	Id             int
	UniqueKey      *string  `json:"unique_key,omitempty"`
	Name           *string  `json:"name,omitempty"`
	Description    *string  `json:"description,omitempty"`
	Price          *float64 `json:"price,omitempty"`
	RemainingStock *int     `json:"remaining_stock,omitempty"`
	Rating         *float64 `json:"rating,omitempty"`
}
