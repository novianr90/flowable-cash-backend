package models

import "time"

type Daily struct {
	ID          uint      `gorm:"not null;primaryKey"`
	Date        time.Time `gorm:"not null"`
	Name        string    `gorm:"not null"`
	Type        string
	Total       uint `gorm:"not null"`
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
