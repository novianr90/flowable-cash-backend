package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Purchase struct {
	ID            uint      `gorm:"not null;primaryKey"`
	Date          time.Time `gorm:"not null"`
	Name          string    `gorm:"not null"`
	Total         uint      `gorm:"not null"`
	TransactionID uint      `gorm:"not null"`
	Description   string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (s *Purchase) BeforeCreate(tx *gorm.DB) error {
	var sale Sale

	err := tx.Where("transaction_id = ?", sale.TransactionID).First(&sale).Error

	if err == nil {
		return errors.New("the transactions already added to sale")
	}

	return nil
}
