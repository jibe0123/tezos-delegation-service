package config

import (
	"os"
)

// Config holds the configuration values.
type Config struct {
	TezosAPIBaseURL string
	DatabasePath    string
}

// LoadConfig loads configuration from environment variables.
func LoadConfig() Config {
	config := Config{
		TezosAPIBaseURL: getEnv("TEZOS_API_BASE_URL", "https://api.tzkt.io/v1/"),
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
