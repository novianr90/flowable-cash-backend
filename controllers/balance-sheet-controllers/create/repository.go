package createBalanceSheet

import (
	"errors"
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

	var found bool = false

	model := r.db.Model(&models.BalanceSheet{})

	balanceSheet := models.BalanceSheet{
		AccountNo:   input.AccountNo,
		AccountName: input.AccountName,
		Balance:     input.Balance,
	}

	if err := model.Raw("SELECT EXIST(SELECT 1 FROM balance_sheets WHERE account_name = ?)", input.AccountName).
		Scan(&found).Error; err != nil {
		return &models.BalanceSheet{}, err
	}

	if found {
		return &models.BalanceSheet{}, errors.New("data already created")
	}

	if err := model.FirstOrCreate(&balanceSheet).Error; err != nil {
		return &models.BalanceSheet{}, err
	}

	return &balanceSheet, nil
}
