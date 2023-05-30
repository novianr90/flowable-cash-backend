package models

import "time"

type Sale struct {
	ID           uint      `gorm:"not null;primaryKey"`
	Date         time.Time `gorm:"not null"`
	Name         string    `gorm:"not null"`
	Total        uint      `gorm:"not null"`
	Description  string
	TransctionID uint `gorm:"not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
