package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string
	Port        string
	Environment string
	JWTSecret     string
	JWTExpiration time.Duration
}

func Load() *Config {
	if error := godotenv.Load(); error != nil {
		log.Println("no .env file found, cp .env.dist first.")
	}

	expirationStr := getEnv("JWT_EXPIRATION", "24h")
	expiration, err := time.ParseDuration(expirationStr)
	if err != nil {
		expiration = 168 * time.Hour
	}

	return &Config{
		DatabaseURL:   getEnv("DATABASE_URL", ""),
		Port:          getEnv("PORT", "3000"),
		Environment:   getEnv("GO_ENV", "development"),
		JWTSecret:     getEnv("JWT_SECRET", ""),
		JWTExpiration: expiration,
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}