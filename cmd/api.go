package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5"
	repo "github.com/thethoomm/ecom/backend/internal/adapters/postgresql/sqlc"
	"github.com/thethoomm/ecom/backend/internal/products"
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

	productService := products.NewService(repo.New(api.db))
	productHandler := products.NewHandler(productService)
	router.Get("/products", productHandler.ListProducts)
	router.Get("/products/{id}", productHandler.FindProductById)

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
