package models

import (
	"time"

	"gorm.io/gorm"
)

type Posting struct {
	ID            uint      `gorm:"primaryKey; not null" json:"posting_id" form:"posting_id"`
	TransactionID uint      `gorm:"not null" json:"transaction_id" form:"transaction_id"`
	AccountID     uint      `gorm:"not null" json:"account_id" form:"account_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func (p *Posting) BeforeCreate(tx *gorm.DB) error {
	if isDuplicateID(tx, p.TransactionID) {
		return gorm.ErrDuplicatedKey
	}

	return nil
}

func isDuplicateID(tx *gorm.DB, id uint) bool {
	var count int64
	tx.Model(&Posting{}).Where("transaction_id = ?", id).Count(&count)
	return count > 0
}
