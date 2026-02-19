package products

import (
	"context"

	repo "github.com/thethoomm/ecom/backend/internal/adapters/postgresql/sqlc"
)

type ProductsService interface {
	ListProducts(ctx context.Context) ([]repo.Product, error)
	FindProductById(ctx context.Context, id int64) (repo.Product, error)
}

type productsService struct {
	repo repo.Querier
}

func NewProductsService(repo repo.Querier) ProductsService {
	return &productsService{
		repo: repo,
	}
}

func (s *productsService) ListProducts(ctx context.Context) ([]repo.Product, error) {
	products, err := s.repo.ListProducts(ctx)

	return products, err
}

func (s *productsService) FindProductById(ctx context.Context, id int64) (repo.Product, error) {
	product, err := s.repo.FindProductById(ctx, id)

	return product, err
}
