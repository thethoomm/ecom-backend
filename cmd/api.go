package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5"
	repo "github.com/thethoomm/ecom/backend/internal/adapters/postgresql/sqlc"
	"github.com/thethoomm/ecom/backend/internal/products"
	"github.com/thethoomm/ecom/backend/internal/users"
	"go.uber.org/zap"
)

func (api *api) mount() http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	// If processing takes longer than 60s, it will be stopped
	router.Use(middleware.Timeout(time.Second * 60))

	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("i am good"))
	})

	repository := repo.New(api.db)

	productService := products.NewProductsService(repository)
	productHandler := products.NewProductsHandler(productService)
	router.Route("/products", func(r chi.Router) {
		r.Get("/", productHandler.ListProducts)
		r.Get("/{id}", productHandler.FindProductById)
	})

	usersService := users.NewUsersService(repository)
	usersHandler := users.NewUsersHandler(usersService)
	router.Route("/users", func(r chi.Router) {
		r.Post("/", usersHandler.CreateUser)
	})

	return router
}

func (api *api) run(h http.Handler) error {
	server := &http.Server{
		Addr:         api.config.addr,
		Handler:      h,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	zap.S().Infof("server has been started at add %s", api.config.addr)

	return server.ListenAndServe()
}

type api struct {
	config config
	db     *pgx.Conn
}

type config struct {
	addr string
	db   dbConfig
}

type dbConfig struct {
	url string
}
