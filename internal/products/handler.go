package products

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/thethoomm/ecom/backend/internal/json"
	"go.uber.org/zap"
)

type productsHandler struct {
	service ProductsService
}

func NewProductsHandler(service ProductsService) *productsHandler {
	return &productsHandler{
		service: service,
	}
}

func (h *productsHandler) ListProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.service.ListProducts(r.Context())

	if err != nil {
		zap.S().Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.Write(w, http.StatusOK, products)
}

func (h *productsHandler) FindProductById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		zap.S().Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	product, err := h.service.FindProductById(r.Context(), id)
	if err != nil {
		zap.S().Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.Write(w, http.StatusOK, product)
}
