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

	var check models.BalanceSheet

	model := r.db.Model(&models.BalanceSheet{})

	balanceSheet := models.BalanceSheet{
		AccountNo:   input.AccountNo,
		AccountName: input.AccountName,
		Balance:     input.Balance,
	}

	if isExist := model.Where("account_name = ?", input.AccountName).First(&check).RowsAffected; isExist > 0 {
		return &models.BalanceSheet{}, errors.New("data already created")
	}

	if err := model.Create(&balanceSheet).Error; err != nil {
		return &models.BalanceSheet{}, err
	}

	return &balanceSheet, nil
}
