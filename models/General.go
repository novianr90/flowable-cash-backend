package models

import "time"

type Journal struct {
	ID          uint      `gorm:"not null;primaryKey"`
	Date        time.Time `gorm:"not null"`
	Description string
	Debit       float64
	Credit      float64
	TotalDebit  float64
	TotalCredit float64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
