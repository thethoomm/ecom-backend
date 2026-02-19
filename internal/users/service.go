package users

import (
	"context"
	"net/mail"

	repo "github.com/thethoomm/ecom/backend/internal/adapters/postgresql/sqlc"
	"github.com/thethoomm/ecom/backend/internal/utils"
)

type UsersService interface {
	CreateUser(ctx context.Context, createUserParams repo.CreateUserParams) (repo.CreateUserRow, error)
}

type usersService struct {
	repo repo.Querier
}

func NewUsersService(repo repo.Querier) UsersService {
	return &usersService{
		repo: repo,
	}
}

func (s *usersService) CreateUser(ctx context.Context, createUserParams repo.CreateUserParams) (repo.CreateUserRow, error) {
	var insertData repo.CreateUserParams

	_, err := mail.ParseAddress(createUserParams.Email)
	if err != nil {
		return repo.CreateUserRow{}, err
	}
	insertData.Email = createUserParams.Email

	hashedPassword, err := utils.HashPassword(createUserParams.Password)
	if err != nil {
		return repo.CreateUserRow{}, err
	}
	insertData.Password = hashedPassword

	insertData.Name = createUserParams.Name
	user, err := s.repo.CreateUser(ctx, insertData)

	return user, err
}
