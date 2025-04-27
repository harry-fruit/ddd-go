package router

import (
	"net/http"

	httphandler "github.com/harry-fruit/ddd-go/internal/infrastructure/api/http/handler"
)

func RegisterProductRoutes(mux *http.ServeMux, productHandler *httphandler.ProductHandler) {
	mux.HandleFunc("/products", productHandler.GetEntryPoints)
}
