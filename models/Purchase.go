package models

import "time"

type Purchase struct {
	ID           uint      `gorm:"not null;primaryKey"`
	Date         time.Time `gorm:"not null"`
	Name         string    `gorm:"not null"`
	Total        uint      `gorm:"not null"`
	TransctionID uint      `gorm:"not null"`
	Description  string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
