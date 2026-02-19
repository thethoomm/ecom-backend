package main

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/thethoomm/ecom/backend/internal/env"
	"go.uber.org/zap"
)

func main() {
	env.Load()
	ctx := context.Background()

	cfg := config{
		addr: ":8080",
		db: dbConfig{
			url: env.GetString("GOOSE_DBSTRING", "postgres://root:root@localhost:5432/ecom"),
		},
	}

	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)
	defer logger.Sync()

	conn, err := pgx.Connect(ctx, cfg.db.url)
	if err != nil {
		panic(err)
	}
	defer conn.Close(ctx)

	logger.Sugar().Infof("connected to database: %s", cfg.db.url)

	api := api{
		config: cfg,
		db:     conn,
	}

	if err := api.run(api.mount()); err != nil {
		zap.S().Error("server has failed to start, err: ", zap.Error(err))
		os.Exit(1)
	}
}
