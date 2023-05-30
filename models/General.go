package models

import "time"

type General struct {
	ID            uint      `gorm:"not null;primaryKey"`
	Date          time.Time `gorm:"not null"`
	Description   string
	Balance       Balance
	TransactionID uint
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
