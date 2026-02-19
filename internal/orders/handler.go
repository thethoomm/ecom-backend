package orders

import (
	"net/http"

	"github.com/thethoomm/ecom/backend/internal/json"
	"go.uber.org/zap"
)

type ordersHandler struct {
	service OrdersService
}

func NewOrdersHandler(service OrdersService) *ordersHandler {
	return &ordersHandler{
		service: service,
	}
}

func (h *ordersHandler) PlaceOrder(w http.ResponseWriter, r *http.Request) {
	var createOrderWithItemsParams CreateOrderWithItemsParams

	err := json.ParseBody(w, r, &createOrderWithItemsParams)
	if err != nil {
		zap.S().Error(err.Error())

		if err == json.ErrEmptyBody {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	createdOrder, err := h.service.PlaceOrder(r.Context(), createOrderWithItemsParams)
	if err != nil {
		zap.S().Error(err.Error())

		if err == ErrProductNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.Write(w, http.StatusCreated, createdOrder)
}
