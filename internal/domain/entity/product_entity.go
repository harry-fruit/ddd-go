package entity

type Product struct {
	Id             int
	UniqueKey      string
	Name           string
	Description    string
	Price          float64
	RemainingStock int
	Rating         float64
}

func NewProduct(uniqueKey string, name string, description string, price float64, remainingStock int, rating float64) *Product {
	return &Product{
		UniqueKey:      uniqueKey,
		Name:           name,
		Description:    description,
		Price:          price,
		RemainingStock: remainingStock,
		Rating:         rating,
	}
}
