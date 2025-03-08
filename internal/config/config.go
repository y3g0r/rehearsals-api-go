package config

import (
	"os"
)

type Config struct {
	ServerAddress string
	SecretKey     string
}

func LoadConfig() *Config {
	return &Config{
		ServerAddress: getEnv("SERVER_ADDRESS", ":8000"),
		SecretKey:     getOrPanic("SECRET_KEY"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getOrPanic(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	panic("missing required environment variable: " + key)
}
