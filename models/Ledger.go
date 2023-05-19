package models

import "time"

type Ledger struct {
	ID          uint `gorm:"not null;primaryKey"`
	Account     string
	Date        time.Time
	Description string
	Balance     Balance
	GeneralReff uint
	General     General `gorm:"foreignKey:GeneralReff"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
