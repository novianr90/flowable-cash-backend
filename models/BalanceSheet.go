package models

import (
	"errors"
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
	var balanceSheet BalanceSheet

	accountNameToCheck := []string{
		"Kas",
		"Persediaan Barang Dagang",
		"Device",
		"Hutang Dagang",
		"Modal",
		"Laba Disimpan",
		"Mengambil Laba",
		"Penjualan",
		"Pembelian",
		"Beban Pembelian",
		"Beban Penjualan",
	}

	for _, value := range accountNameToCheck {
		err := tx.Where("account_name = ?", value).First(&balanceSheet).Error

		if err == nil {
			return errors.New("data already filled, please delete or update")
		}
	}

	_ = balanceSheet

	return nil

}
