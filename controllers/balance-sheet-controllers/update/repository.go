package updateBalanceSheet

import (
	"errors"
	"flowable-cash-backend/models"

	"gorm.io/gorm"
)

type Repository interface {
	UpdateBalanceSheet(input *models.BalanceSheet) (*models.BalanceSheet, error)
}

type repository struct {
	db *gorm.DB
}

func NewUpdateBalanceSheetRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) UpdateBalanceSheet(input *models.BalanceSheet) (*models.BalanceSheet, error) {
	model := r.db.Model(&models.BalanceSheet{})

	query := models.BalanceSheet{
		Balance: input.Balance,
	}

	res := model.Where("account_name = ?", input.AccountName).Updates(&query)

	if res.RowsAffected == 0 {
		return &models.BalanceSheet{}, errors.New("no record to update")
	}

	if res.Error != nil {
		return &models.BalanceSheet{}, res.Error
	}

	return &query, nil
}
