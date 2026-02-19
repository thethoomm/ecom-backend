package orders

import repo "github.com/thethoomm/ecom/backend/internal/adapters/postgresql/sqlc"

type CreateOrderWithItemsParams struct {
	CustomerID int64                  `json:"customer_id"`
	Items      []CreateOrderItemInput `json:"items"`
}

type CreateOrderItemInput struct {
	ProductID int64 `json:"product_id"`
	Quantity  int32 `json:"quantity"`
}

type OrderResponse struct {
	repo.Order
	Items []OrderItemReponse `json:"items"`
}

type OrderItemReponse struct {
	ID       int64 `json:"id"`
	Price    int32 `json:"price"`
	Quantity int32 `json:"quantity"`
}
