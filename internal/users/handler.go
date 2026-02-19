package users

import (
	"net/http"

	repo "github.com/thethoomm/ecom/backend/internal/adapters/postgresql/sqlc"
	"github.com/thethoomm/ecom/backend/internal/json"
	"go.uber.org/zap"
)

type usersHandler struct {
	service UsersService
}

func NewUsersHandler(service UsersService) *usersHandler {
	return &usersHandler{
		service: service,
	}
}

func (h *usersHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var createUserParams repo.CreateUserParams

	err := json.ParseBody(w, r, &createUserParams)
	if err != nil {
		zap.S().Error(err.Error())
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	createdUser, err := h.service.CreateUser(r.Context(), createUserParams)
	if err != nil {
		zap.S().Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.Write(w, http.StatusCreated, createdUser)
}
