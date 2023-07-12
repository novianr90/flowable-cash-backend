package models

import "time"

type Posting struct {
	ID            uint      `gorm:"primaryKey; not null" json:"posting_id" form:"posting_id"`
	TransactionID uint      `gorm:"not null" json:"transaction_id" form:"transaction_id"`
	AccountID     uint      `gorm:"not null" json:"account_id" form:"account_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
