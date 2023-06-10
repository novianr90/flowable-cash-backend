package models

import "time"

type BalanceSheet struct {
	ID          uint    `gorm:"primaryKey;not null"`
	AccountNo   string  `gorm:"not null"`
	AccountName string  `gorm:"not null"`
	Balance     Balance `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
