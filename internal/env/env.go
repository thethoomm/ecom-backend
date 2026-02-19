package env

import (
	"os"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func Load() {
	err := godotenv.Load()
	if err != nil {
		zap.S().Fatal("error loading .env file")
	}
}

func GetString(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}

	return fallback
}
