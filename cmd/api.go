package main

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
