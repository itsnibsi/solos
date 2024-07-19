package config

import (
	"os"
	"strconv"
)

type Config struct {
	Bind string
	Port int
	Env  string
}

func Load() *Config {
	return &Config{
		Bind: getEnv("BIND", "0.0.0.0"),
		Port: getEnvAsInt("PORT", 8080),
		Env:  getEnv("ENV", "development"),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func getEnvAsInt(key string, fallback int) int {
	strValue := getEnv(key, "")
	if value, err := strconv.Atoi(strValue); err == nil {
		return value
	}
	return fallback
}
