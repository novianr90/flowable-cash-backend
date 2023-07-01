package updateBalanceSheet

import (
	"encoding/json"
	"errors"
	"flowable-cash-backend/models"

	"gorm.io/gorm"
)

type Repository interface {
	UpdateBalanceSheet(input *models.BalanceSheet) (*models.BalanceSheet, uint, error)
}

type repository struct {
	db *gorm.DB
}

func NewUpdateBalanceSheetRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) UpdateBalanceSheet(input *models.BalanceSheet) (*models.BalanceSheet, uint, error) {
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
		return &models.BalanceSheet{}, 2, errors.New("no record to update")
	}

	if res.Error != nil {
		return &models.BalanceSheet{}, 2, res.Error
	}

	_ = model.Where("account_name = ?", input.AccountName).First(&response)

	var balanceResponse models.Balance

	_ = json.Unmarshal(response.Balance, &balanceResponse)

	if (newBalance.Debit == balanceResponse.Debit) && (newBalance.Credit == balanceResponse.Credit) {
		return &response, 0, nil
	}

	return &response, 1, nil
}
