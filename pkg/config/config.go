package config

import (
	"os"
)

// Config holds the configuration values.
type Config struct {
	TezosAPIBaseURL string
	DBHost          string
	DBPort          string
	DBUser          string
	DBPassword      string
	DBName          string
}

// LoadConfig loads configuration from environment variables.
func LoadConfig() Config {
	config := Config{
		TezosAPIBaseURL: getEnv("TZKT_API_BASE_URL", "https://api.tzkt.io/v1/"),
		DBHost:          getEnv("DB_HOST", "localhost"),
		DBPort:          getEnv("DB_PORT", "3306"),
		DBUser:          getEnv("DB_USER", "root"),
		DBPassword:      getEnv("DB_PASSWORD", "root"),
		DBName:          getEnv("DB_NAME", "tezos_delegations"),
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
