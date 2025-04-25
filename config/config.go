package config

import (
	"os"
)

// Config holds all configuration for the application
type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

// LoadConfig loads configuration from environment variables with default values
func LoadConfig() *Config {
	return &Config{
		DBHost:     getEnvOrDefault("DB_HOST", "localhost"),
		DBPort:     getEnvOrDefault("DB_PORT", "3306"),
		DBUser:     getEnvOrDefault("DB_USER", "root"),
		DBPassword: getEnvOrDefault("DB_PASSWORD", "password"),
		DBName:     getEnvOrDefault("DB_NAME", "my_db"),
	}
}

// getEnvOrDefault gets an environment variable or returns the default value if not set
func getEnvOrDefault(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
