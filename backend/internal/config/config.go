package config

import (
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Database DatabaseConfig
	Server   ServerConfig
	JWT      JWTConfig
}

type DatabaseConfig struct {
	URL string
}

type ServerConfig struct {
	Port    string
	GinMode string
}

type JWTConfig struct {
	Secret                    string
	ExpirationHours           int
	RefreshTokenExpirationDays int
}

func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		// .env файл не обязателен, если переменные заданы через окружение
	}

	return &Config{
		Database: DatabaseConfig{
			URL: getEnv("DATABASE_URL", "postgres://postgres:postgres@localhost:51213/trello_db?sslmode=disable"),
		},
		Server: ServerConfig{
			Port:    getEnv("PORT", "8080"),
			GinMode: getEnv("GIN_MODE", "debug"),
		},
		JWT: JWTConfig{
			Secret:                     getEnv("JWT_SECRET", "your-secret-key-change-in-production"),
			ExpirationHours:            getEnvAsInt("JWT_EXPIRATION_HOURS", 24),
			RefreshTokenExpirationDays: getEnvAsInt("REFRESH_TOKEN_EXPIRATION_DAYS", 7),
		},
	}, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	valueStr := os.Getenv(key)
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultValue
}

func (c *JWTConfig) GetAccessTokenExpiration() time.Duration {
	return time.Duration(c.ExpirationHours) * time.Hour
}

func (c *JWTConfig) GetRefreshTokenExpiration() time.Duration {
	return time.Duration(c.RefreshTokenExpirationDays) * 24 * time.Hour
}

