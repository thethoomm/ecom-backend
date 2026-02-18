package main

import (
	"os"

	"go.uber.org/zap"
)

func main() {
	cfg := config{
		addr: ":8080",
		db:   dbConfig{},
	}

	api := api{
		config: cfg,
	}

	logger, _ := zap.NewDevelopment()
	defer logger.Sync()

	zap.ReplaceGlobals(logger)

	if err := api.run(api.mount()); err != nil {
		zap.S().Error("server has failed to start, err: ", zap.Error(err))
		os.Exit(1)
	}
}
