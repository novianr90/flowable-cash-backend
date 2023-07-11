package models

import "time"

type Transaction struct {
	ID            uint      `gorm:"not null;primaryKey"`
	Date          time.Time `gorm:"not null"`
	Name          string    `gorm:"not null"`
	Type          string
	Payment       string `gorm:"not null"`
	Total         uint   `gorm:"not null"`
	FeeType       string
	Fee           uint
	Description   string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	AlreadyPosted uint
}
