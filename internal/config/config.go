package config

import "os"

// Config holds all configuration for the application.
type Config struct {
	Port        string
	DatabaseURL string
	RedisURL    string
	BaseURL     string
}

// Load reads configuration from environment variables with sensible defaults.
func Load() *Config {
	return &Config{
		Port:        getEnv("PORT", "8080"),
		DatabaseURL: getEnv("DATABASE_URL", "postgres://user:pass@localhost:5432/shortener?sslmode=disable"),
		RedisURL:    getEnv("REDIS_URL", "localhost:6379"),
		BaseURL:     getEnv("BASE_URL", "http://localhost:8080"),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
