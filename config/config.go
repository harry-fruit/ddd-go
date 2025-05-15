package config

import (
	"flag"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort  string
	DatabaseURL string
	Environment string
	LogLevel    string
	DbHost      string
	DbPort      string
	DbUser      string
	DbPassword  string
	DbName      string
	DbSslmode   string
}

func LoadConfig() (*Config, error) {
	env, err := getRunMode()
	if err != nil {
		return nil, err
	}

	fileName := getEnvFileName(*env)

	if err := godotenv.Load(fileName); err != nil {
		return nil, fmt.Errorf("error loading .env file: %v", err)
	}

	return &Config{
		Environment: *env,
		ServerPort:  getEnv("API_PORT", "3001"),
		LogLevel:    getEnv("LOG_LEVEL", "info"),
		DbHost:      getEnv("DB_HOST", ""),
		DbPort:      getEnv("DB_PORT", ""),
		DbUser:      getEnv("DB_USER", ""),
		DbPassword:  getEnv("DB_PASSWORD", ""),
		DbName:      getEnv("DB_NAME", ""),
		DbSslmode:   getEnv("DB_SSLMODE", ""),
	}, nil
}

func getEnv(key string, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}

func getRunMode() (result *string, err error) {
	envFlag := flag.String("env", "", "Environment")

	flag.Parse()

	env := *envFlag
	if env == "" {
		return nil, fmt.Errorf("no environment specified")
	}

	return &env, nil
}

func getEnvFileName(env string) string {
	return fmt.Sprintf(".env.%s", env)
}
