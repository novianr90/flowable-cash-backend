package models

import "time"

type Accounts struct {
	AccountID     uint      `gorm:"primaryKey;not null"`
	AccountDate   time.Time `gorm:"not null"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	BalanceSheets []BalanceSheet `gorm:"foreignKey:AccountID"`
}
