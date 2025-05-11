package entity

type Product struct {
	ID             uint
	UniqueKey      string
	Name           string
	Description    *string
	Price          float64
	RemainingStock *int
}

func NewProduct(uniqueKey string, name string, description *string, price float64, remainingStock *int, rating float64) *Product {
	return &Product{
		UniqueKey:      uniqueKey,
		Name:           name,
		Description:    description,
		Price:          price,
		RemainingStock: remainingStock,
	}
}
