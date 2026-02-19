package products

import (
	"net/http"

	"github.com/thethoomm/ecom/backend/internal/json"
	"go.uber.org/zap"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) ListProducts(w http.ResponseWriter, r *http.Request) {
	err := h.service.ListProducts(r.Context())

	if err != nil {
		zap.S().Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	products := []string{}

	json.Write(w, http.StatusOK, products)

}
