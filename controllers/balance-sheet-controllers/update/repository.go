package updateBalanceSheet

import (
	"encoding/json"
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

	var response models.BalanceSheet

	var inputBalance models.Balance

	var localBalance models.Balance

	_ = json.Unmarshal(input.Balance, &inputBalance)

	_ = model.Where("account_name = ?", input.AccountName).First(&response)

	_ = json.Unmarshal(response.Balance, &localBalance)

	newBalance := models.Balance{
		Debit:  inputBalance.Debit + localBalance.Debit,
		Credit: inputBalance.Credit + localBalance.Credit,
	}

	formattedBalance, _ := json.Marshal(&newBalance)

	query := models.BalanceSheet{
		Balance: formattedBalance,
	}

	res := model.Where("account_name = ?", input.AccountName).Updates(&query)

	if res.RowsAffected == 0 {
		return &models.BalanceSheet{}, errors.New("no record to update")
	}

	if res.Error != nil {
		return &models.BalanceSheet{}, res.Error
	}

	_ = model.Where("account_name = ?", input.AccountName).First(&response)

	return &response, nil
}
