package configs

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	Host                   string
	Port                   string
	User                   string
	Password               string
	Address                string
	Name                   string
	JWTSecret              string
	JWTExpirationInSeconds int64
}

var Envs = initializeConfig()

func initializeConfig() DBConfig {
	godotenv.Load()

	return DBConfig{
		Host:                   getEnv("PUBLIC_HOST", "http://localhost"),
		Port:                   getEnv("PORT", "4001"),
		User:                   getEnv("DB_USER", "root"),
		Password:               getEnv("DB_PASSWORD", "password"),
		Address:                fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "3306")),
		Name:                   getEnv("DB_NAME", "gostore-api"),
		JWTSecret:              getEnv("JWT_SECRET", "not_a_secret_but_it_should"),
		JWTExpirationInSeconds: getEnvAsInt("JWT_EXPIRATION_IN_SECONDS", 3600*24*7),
	}
}

func getEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func getEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}

		return i
	}

	return fallback
}
