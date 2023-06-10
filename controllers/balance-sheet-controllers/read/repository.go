package readBalanceSheet

import (
	"flowable-cash-backend/models"

	"gorm.io/gorm"
)

type Repository interface {
	GetBalanceSheet() (*[]models.BalanceSheet, error)
}

type repository struct {
	db *gorm.DB
}

func NewReadBalanceSheetRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) GetBalanceSheet() (*[]models.BalanceSheet, error) {
	var balanceSheet []models.BalanceSheet

	model := r.db.Model(&models.BalanceSheet{})

	err := model.Find(&balanceSheet).Error

	if err != nil {
		return &[]models.BalanceSheet{}, err
	}

	return &balanceSheet, nil
}
