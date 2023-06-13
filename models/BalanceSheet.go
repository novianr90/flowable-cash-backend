package models

import (
	"time"

	"gorm.io/gorm"
)

type BalanceSheet struct {
	ID          uint   `gorm:"primaryKey;not null"`
	AccountNo   string `gorm:"not null"`
	AccountName string `gorm:"not null"`
	Balance     []byte `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (b *BalanceSheet) BeforeCreate(tx *gorm.DB) error {

	var balance BalanceSheet

	if err := tx.Model(&b).Where("account_name = ?", b.AccountName).First(&balance).Error; err != nil {
		return err
	}

	return nil
}
