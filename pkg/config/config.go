package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config holds the configuration values.
type Config struct {
	TezosAPIBaseURL string
	DatabasePath    string
}

// LoadConfig loads configuration from environment variables.
func LoadConfig() Config {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	config := Config{
		TezosAPIBaseURL: getEnv("TZKT_API_BASE_URL", "https://api.tzkt.io/v1/"),
		DatabasePath:    getEnv("DATABASE_PATH", "delegations.db"),
	}
	return config
}

// getEnv gets an environment variable or returns a default value.
func getEnv(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = defaultValue
	}
	return value
}
