package models

import (
	"time"
)

type BalanceSheet struct {
	ID          uint   `gorm:"primaryKey;not null"`
	Month       uint   `gorm:"not null"`
	AccountNo   string `gorm:"not null"`
	AccountName string `gorm:"not null"`
	Balance     []byte `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
