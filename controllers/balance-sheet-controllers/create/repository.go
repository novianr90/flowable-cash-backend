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

	var createFlag bool = false

	var check models.BalanceSheet

	model := r.db.Model(&models.BalanceSheet{})

	balanceSheet := models.BalanceSheet{
		AccountNo:   input.AccountNo,
		AccountName: input.AccountName,
		Balance:     input.Balance,
	}

	isExist := model.Where("account_name = ?", input.AccountName).First(&check)

	if isExist.Error == gorm.ErrRecordNotFound {
		createFlag = true
	}

	if isExist.RowsAffected > 0 {
		createFlag = true
	}

	if createFlag {
		if err := model.Create(&balanceSheet).Error; err != nil {
			return &models.BalanceSheet{}, err
		}
	}

	if !createFlag {
		return &models.BalanceSheet{}, errors.New("data already created")
	}

	return &balanceSheet, nil
}
