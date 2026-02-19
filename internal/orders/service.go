package orders

import (
	"context"

	"github.com/jackc/pgx/v5"
	repo "github.com/thethoomm/ecom/backend/internal/adapters/postgresql/sqlc"
)

type OrdersService interface {
	PlaceOrder(ctx context.Context, createOrderWithItemsParams CreateOrderWithItemsParams) (OrderResponse, error)
}

type ordersService struct {
	repo *repo.Queries
	db   *pgx.Conn
}

func NewOrdersService(repo *repo.Queries, db *pgx.Conn) OrdersService {
	return &ordersService{
		repo: repo,
		db:   db,
	}
}

func (s *ordersService) PlaceOrder(ctx context.Context, createOrderWithItemsParams CreateOrderWithItemsParams) (OrderResponse, error) {
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return OrderResponse{}, err
	}
	defer tx.Rollback(ctx)

	qtx := s.repo.WithTx(tx)

	var orderParams repo.CreateOrderParams = repo.CreateOrderParams{
		CustomerID: createOrderWithItemsParams.CustomerID,
		Status:     repo.OrderStatusPending,
	}

	order, err := qtx.CreateOrder(ctx, orderParams)
	if err != nil {
		return OrderResponse{}, err
	}

	var items []OrderItemReponse
	for _, item := range createOrderWithItemsParams.Items {
		product, err := qtx.FindProductById(ctx, item.ProductID)
		if err != nil {
			return OrderResponse{}, ErrProductNotFound
		}

		if product.Quantity < item.Quantity {
			return OrderResponse{}, ErrProductNoStock
		}

		var orderItemParams repo.CreateOrderItemParams = repo.CreateOrderItemParams{
			OrderID:   order.ID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     product.Price,
		}

		createdItem, err := qtx.CreateOrderItem(ctx, orderItemParams)
		if err != nil {
			return OrderResponse{}, err
		}
		var orderItemResponse OrderItemReponse = OrderItemReponse{
			ID:       createdItem.ID,
			Price:    createdItem.Price,
			Quantity: createdItem.Quantity,
		}
		items = append(items, orderItemResponse)

		// update the product stock
	}

	if err := tx.Commit(ctx); err != nil {
		return OrderResponse{}, err
	}

	var response OrderResponse = OrderResponse{
		Order: order,
		Items: items,
	}

	return response, nil
}
