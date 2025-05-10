package gormmodel

import (
	"time"

	"gorm.io/gorm"
)

type ProductModel struct {
	gorm.Model
	ID             uint
	UniqueKey      string         `gorm:"column:unique_key;not null;unique;size:100"`
	Name           string         `gorm:"column:name;not null;size:150"`
	Description    *string        `gorm:"column:description;size:500"`
	Price          float64        `gorm:"column:price;not null"`
	RemainingStock int            `gorm:"column:remaining_stock;not null"`
	CreatedAt      time.Time      `gorm:"column:created_at;not null"`
	UpdatedAt      time.Time      `gorm:"column:updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"column:deleted_at"`
}
