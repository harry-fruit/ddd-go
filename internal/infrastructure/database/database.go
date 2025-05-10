package infrastructure

import (
	"fmt"
	"strconv"

	"github.com/harry-fruit/ddd-go/config"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func NewSQLDatabase(config *config.Config) (*gorm.DB, error) {
	dbConfig, err := getDbConfig(config)
	if err != nil {
		return nil, err
	}
	dns := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		dbConfig.Host, dbConfig.User, dbConfig.Password, dbConfig.DBName, dbConfig.Port, dbConfig.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	// Auto-migrate tables
	// db.AutoMigrate(&entity.Product{})

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
