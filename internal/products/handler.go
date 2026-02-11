package products

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) ListProduct(w http.ResponseWriter, r *http.Request) {
	// call the service -> ListProduct
	product := []string{"hellow", "world"}
	json.NewEncoder(w).Encode(product)
}
