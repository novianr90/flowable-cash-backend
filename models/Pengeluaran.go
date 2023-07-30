package models

import "time"

type Pengeluaran struct {
	ID            uint      `gorm:"not null;primaryKey"`
	Date          time.Time `gorm:"not null"`
	Name          string    `gorm:"not null"`
	Payment       string    `gorm:"not null"`
	Total         uint      `gorm:"not null"`
	Description   string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	AlreadyPosted uint
	Purchase      string
}
