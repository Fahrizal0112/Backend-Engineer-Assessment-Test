package config

import (
	"flag"
	"fmt"
	"os"

	"banking-service/pkg/logger"

	"github.com/joho/godotenv"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
}

type ServerConfig struct {
	Host string
	Port int
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		logger.Warning("No. env file found, using environment variables")
	}

	host := flag.String("host", "0.0.0.0", "Host for the server")
	port := flag.Int("port", 8080, "Port for the server")
	flag.Parse()

	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5432")
	dbUser := getEnv("DB_USER", "postgres")
	dbPassword := getEnv("DB_PASSWORD", "")
	dbName := getEnv("DB_NAME", "bank")

	logger.Info(fmt.Sprintf("Loaded configuration - Server: %s:%d, Database: %s:%s/%s",
		*host, *port, dbHost, dbPort, dbName))

	return &Config{
		Server: ServerConfig{
			Host: *host,
			Port: *port,
		},
		Database: DatabaseConfig{
			Host:     dbHost,
			Port:     dbPort,
			User:     dbUser,
			Password: dbPassword,
			Name:     dbName,
		},
	}
}

func getEnv(key, fallback string) string {
	if value, exist := os.LookupEnv(key); exist {
		return value
	}
	return fallback
}
