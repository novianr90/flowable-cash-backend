package readBalanceSheet

import (
	"flowable-cash-backend/models"

	"gorm.io/gorm"
)

type Repository interface {
	GetAllAccounts(input *models.BalanceSheet) (*[]models.BalanceSheet, error)
	GetAllAccountsByAccountName(input *models.BalanceSheet) (*models.BalanceSheet, error)
}

type repository struct {
	db *gorm.DB
}

func NewReadBalanceSheetRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) GetAllAccounts(input *models.BalanceSheet) (*[]models.BalanceSheet, error) {
	var balanceSheet []models.BalanceSheet

	model := r.db.Model(&models.BalanceSheet{})

	err := model.Where("month = ?", input.Month).Find(&balanceSheet).Error

	if err != nil {
		return &[]models.BalanceSheet{}, err
	}

	return &balanceSheet, nil
}

func (r *repository) GetAllAccountsByAccountName(input *models.BalanceSheet) (*models.BalanceSheet, error) {
	var balanceSheet models.BalanceSheet

	model := r.db.Model(&models.BalanceSheet{})

	err := model.
		Where("month = ?", input.Month).
		Where("account_name = ?", input.AccountName).
		First(&balanceSheet).Error

	if err != nil {
		return &models.BalanceSheet{}, err
	}

	return &balanceSheet, nil
}
