package createBalanceSheet

import (
	"flowable-cash-backend/models"

	"gorm.io/gorm"
)

type Repository interface {
	CreateBalanceSheet(input *models.BalanceSheet) (*models.BalanceSheet, error)
}

type repository struct {
	db *gorm.DB
}

func NewCreateBalanceSheetRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) CreateBalanceSheet(input *models.BalanceSheet) (*models.BalanceSheet, error) {

	balanceSheet := models.BalanceSheet{
		AccountNo:   input.AccountNo,
		AccountName: input.AccountName,
		Balance:     input.Balance,
	}

	if err := r.db.Create(&balanceSheet).Error; err != nil {
		return &models.BalanceSheet{}, err
	}

	return &balanceSheet, nil
}
