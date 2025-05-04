package infrastructure

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/harry-fruit/ddd-go/config"
	_ "github.com/lib/pq"
)

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(config *config.Config) (*sql.DB, error) {
	dbConfig, err := getDbConfig(config)
	if err != nil {
		return nil, err
	}

	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.DBName, dbConfig.SSLMode,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open db: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping db: %w", err)
	}

	return db, nil
}

func getDbConfig(config *config.Config) (*DatabaseConfig, error) {
	port, err := strconv.Atoi(config.DbPort)

	if err != nil {
		return nil, fmt.Errorf("failed to parse db port: %w", err)
	}

	return &DatabaseConfig{
		Host:     config.DbHost,
		Port:     port,
		User:     config.DbUser,
		Password: config.DbPassword,
		DBName:   config.DbName,
		SSLMode:  config.DbSslmode,
	}, nil
}
