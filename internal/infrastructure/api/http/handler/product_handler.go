package httphandler

import (
	"net/http"

	"github.com/harry-fruit/ddd-go/internal/domain/service"
)

type ProductHandler struct {
	productService *service.ProductService
}

func NewProductHandler(service *service.ProductService) *ProductHandler {
	return &ProductHandler{
		productService: service,
	}
}

func (ph *ProductHandler) createProduct(w http.ResponseWriter, r *http.Request) {

}

func (ph *ProductHandler) GetEntryPoints(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	// case http.MethodGet:
	//     ph.handleGetProducts(w, r)
	case http.MethodPost:
		ph.createProduct(w, r)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

// func (ph *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {

// }
