package models

import "time"

type Ledger struct {
	ID          uint `gorm:"not null;primaryKey"`
	Account     string
	Date        time.Time
	Description string
	Balance     Balance
}
