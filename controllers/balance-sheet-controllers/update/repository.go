package updateBalanceSheet

import (
	"encoding/json"
	"errors"
	"flowable-cash-backend/models"

	"gorm.io/gorm"
)

type Repository interface {
	UpdateAccount(input *models.BalanceSheet) (*models.BalanceSheet, error)
	UpdateAccountAdmin(input *models.BalanceSheet) error
}

type repository struct {
	db *gorm.DB
}

func NewUpdateBalanceSheetRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) UpdateAccount(input *models.BalanceSheet) (*models.BalanceSheet, error) {
	model := r.db.Model(&models.BalanceSheet{})

	var response models.BalanceSheet

	var inputBalance models.Balance

	var localBalance models.Balance

	_ = json.Unmarshal(input.Balance, &inputBalance)

	err := model.
		Where("month = ?", input.Month).
		Where("account_name = ?", input.AccountName).
		First(&response).Error

	if err != nil {
		return &models.BalanceSheet{}, err
	}

	_ = json.Unmarshal(response.Balance, &localBalance)

	newDebit := inputBalance.Debit - localBalance.Debit
	newCredit := inputBalance.Credit - localBalance.Credit

	newBalance := models.Balance{
		Debit:  localBalance.Debit + newDebit,
		Credit: localBalance.Credit + newCredit,
	}

	formattedBalance, _ := json.Marshal(&newBalance)

	query := models.BalanceSheet{
		Balance: formattedBalance,
	}

	res := model.
		Where("month = ?", input.Month).
		Where("account_name = ?", input.AccountName).
		Updates(&query)

	if res.RowsAffected == 0 {
		return &models.BalanceSheet{}, errors.New("no record to update")
	}

	if res.Error != nil {
		return &models.BalanceSheet{}, res.Error
	}

	_ = model.
		Where("month = ?", input.Month).
		Where("account_name = ?", input.AccountName).
		First(&response)

	var balanceResponse models.Balance

	_ = json.Unmarshal(response.Balance, &balanceResponse)

	return &response, nil
}

func (r *repository) UpdateAccountAdmin(input *models.BalanceSheet) error {
	model := r.db.Model(&models.BalanceSheet{})

	query := models.BalanceSheet{
		Balance: input.Balance,
	}

	allAccountName := []string{
		"Kas",
		"Persediaan Barang Dagang",
		"Perlengkapan",
		"Akumulasi Penyusutan Perlengkapan",
		"Hutang Dagang",
		"Modal",
		"Laba Disimpan",
		"Mengambil Laba",
		"Penjualan",
		"Beban Pembelian",
		"Beban Penjualan",
		"Beban Penyusutan",
		"Beban Perlengkapan",
	}

	for _, value := range allAccountName {
		if err := model.Where("month = ?", input.Month).Where("account_name = ?", value).Updates(&query).Error; err != nil {
			return err
		}
	}

	return nil
}
