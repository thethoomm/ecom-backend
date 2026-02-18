package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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

	log.Printf("server has been started at add %s", api.config.addr)

	return server.ListenAndServe()
}

type api struct {
	config config
}

type config struct {
	addr string
	db   dbConfig
}

type dbConfig struct {
	dsn string
}
