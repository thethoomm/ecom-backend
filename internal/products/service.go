package products

import (
	"context"

	repo "github.com/thethoomm/ecom/backend/internal/adapters/postgresql/sqlc"
)

type Service interface {
	ListProducts(ctx context.Context) ([]repo.Product, error)
	FindProductById(ctx context.Context, id int64) (repo.Product, error)
}

type service struct {
	repo repo.Querier
}

func NewService(repo repo.Querier) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) ListProducts(ctx context.Context) ([]repo.Product, error) {
	products, err := s.repo.ListProducts(ctx)

	return products, err
}

func (s *service) FindProductById(ctx context.Context, id int64) (repo.Product, error) {
	product, err := s.repo.FindProductById(ctx, id)

	return product, err
}
