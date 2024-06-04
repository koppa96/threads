package deps

import (
	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
	"go.uber.org/fx"
	"os"
)

type Config struct {
	DBHost          string
	DBPort          string
	DBName          string
	DBUser          string
	DBPass          string
	BindPort        string
	TemporalAddress string
}

func NewConfig(_ fx.Lifecycle) *Config {
	err := godotenv.Load()
	if err != nil {
		log.New("echo").Warnf("Failed to load .env file: %s", err.Error())
	}

	return &Config{
		DBHost:          os.Getenv("DB_HOST"),
		DBPort:          os.Getenv("DB_PORT"),
		DBName:          os.Getenv("DB_NAME"),
		DBUser:          os.Getenv("DB_USER"),
		DBPass:          os.Getenv("DB_PASS"),
		BindPort:        GetEnvOr("BIND_PORT", "8080"),
		TemporalAddress: GetEnvOr("TEMPORAL_ADDRESS", "localhost:7233"),
	}
}

func GetEnvOr(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	return value
}
