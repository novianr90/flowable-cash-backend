package usecase

import (
	"encoding/json"
	"flowable-cash-backend/models"
	"fmt"
	"log"
	"time"

	"gorm.io/gorm"
)

type UseCase interface {
	PostingPenjualan() error
}

type usecase struct {
	db *gorm.DB
}

func NewUseCaseService(db *gorm.DB) *usecase {
	return &usecase{db: db}
}

func (u *usecase) PostingPenjualan() error {

	accountModel := u.db.Model(&models.BalanceSheet{}).Where("month = ?", time.Now().Month())

	var transactions []models.Transaction
	var kas models.BalanceSheet
	var penjualan models.BalanceSheet
	var piutang models.BalanceSheet

	if err := u.db.
		Where("already_posted = ?", 0).
		Where("type = ?", "Penjualan").
		Find(&transactions).Error; err != nil {
		return err
	}

	if err := accountModel.Where("account_name = ?", "Kas").First(&kas).Error; err != nil {
		return err
	}

	if err := accountModel.Where("account_name = ?", "Penjualan").First(&penjualan).Error; err != nil {
		return err
	}

	if err := accountModel.Where("account_name = ?", "Piutang Dagang").First(&piutang).Error; err != nil {
		return err
	}

	totalPenjualanCash := 0.0
	totalPenjualanNonCash := 0.0

	for _, value := range transactions {

		if value.Payment == "Tunai" {
			totalPenjualanCash += float64(value.Total)
		}

		if value.Payment == "Non-Tunai" {
			totalPenjualanNonCash += float64(value.Total)
		}
	}

	totalPenjualan := totalPenjualanCash + totalPenjualanNonCash

	var kasBalanceInDb models.Balance
	var penjualanBalanceInDb models.Balance
	var piutangBalanceInDb models.Balance

	err := json.Unmarshal(kas.Balance, &kasBalanceInDb)

	if err != nil {
		return err
	}

	err = json.Unmarshal(penjualan.Balance, &penjualanBalanceInDb)

	if err != nil {
		return err
	}

	err = json.Unmarshal(piutang.Balance, &piutangBalanceInDb)

	if err != nil {
		return err
	}

	newKas := models.Balance{
		Debit:  kasBalanceInDb.Debit + totalPenjualanCash,
		Credit: kasBalanceInDb.Credit,
	}

	newPenjualan := models.Balance{
		Debit:  penjualanBalanceInDb.Debit,
		Credit: penjualanBalanceInDb.Credit + totalPenjualan,
	}

	newPiutang := models.Balance{
		Debit:  piutangBalanceInDb.Debit + totalPenjualanNonCash,
		Credit: piutangBalanceInDb.Credit,
	}

	kasFormatted, err := json.Marshal(&newKas)

	if err != nil {
		return err
	}

	penjualanFormatted, err := json.Marshal(&newPenjualan)

	if err != nil {
		return err
	}

	piutangFormatted, err := json.Marshal(&newPiutang)

	if err != nil {
		return err
	}

	kas.Balance = kasFormatted
	penjualan.Balance = penjualanFormatted
	piutang.Balance = piutangFormatted

	if err := accountModel.Save(&kas).Error; err != nil {
		return err
	}

	if err := accountModel.Save(&penjualan).Error; err != nil {
		return err
	}

	if err := accountModel.Save(&piutang).Error; err != nil {
		return err
	}

	for i := range transactions {
		transactions[i].AlreadyPosted = 1
		if err := u.db.Model(&transactions[i]).Update("already_posted", 1).Error; err != nil {
			log.Println("Error Updating Transactions:", err)
			continue
		}
	}

	fmt.Println("Posting Completed!")
	return nil
}
