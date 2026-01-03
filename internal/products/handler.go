package products

import (
	"log"
	"net/http"

	"github.com/asmit990/golangAPI/internal/json"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) ListProducts(w http.ResponseWriter, r *http.Request) {
	// call the service -> ListProduct
	err := h.service.ListProducts(r.Context())
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	product := struct {
		Products []string `json:"products"`
	}{}
	// return JSON in as HTTP response
	json.Write(w, http.StatusOK, product)
}
