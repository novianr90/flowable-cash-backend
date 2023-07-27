package usecase

import (
	"encoding/json"
	"errors"
	"flowable-cash-backend/models"
	"fmt"
	"log"
	"time"

	"gorm.io/gorm"
)

type UseCase interface {
	PostingPemasukan() error
}

type usecase struct {
	db *gorm.DB
}

func NewUseCaseService(db *gorm.DB) *usecase {
	return &usecase{db: db}
}

func (u *usecase) PostingPenjualan() error {

	var transactions []models.Pemasukkan

	if err := u.db.
		Where("already_posted = ?", 0).
		Find(&transactions).Error; err != nil {
		return err
	}

	if transactions == nil {
		return errors.New("there's no transactions to posting at the time")
	}

	var kas models.BalanceSheet
	if err := u.db.
		Where("month = ?", time.Now().Month()).
		Where("account_name = ?", "Kas").
		First(&kas).Error; err != nil {
		return err
	}

	var penjualan models.BalanceSheet
	if err := u.db.
		Where("month = ?", time.Now().Month()).
		Where("account_name = ?", "Penjualan").
		First(&penjualan).Error; err != nil {
		return err
	}

	var piutang models.BalanceSheet
	if err := u.db.
		Where("month = ?", time.Now().Month()).
		Where("account_name = ?", "Piutang Dagang").
		First(&piutang).Error; err != nil {
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

	if err := u.db.Save(&kas).Error; err != nil {
		return err
	}

	if err := u.db.Save(&penjualan).Error; err != nil {
		return err
	}

	if err := u.db.Save(&piutang).Error; err != nil {
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

// func (u *usecase) PostingPengeluaran() error {
// 	var transactions []models.Pengeluaran

// 	if err := u.db.
// 		Where("already_posted = ?", 0).
// 		Find(&transactions).Error; err != nil {
// 		return err
// 	}

// 	var kas models.BalanceSheet
// 	if err := u.db.
// 		Where("month = ?", time.Now().Month()).
// 		Where("account_name = ?", "Kas").
// 		First(&kas).Error; err != nil {
// 		return err
// 	}

// 	var pembelian models.BalanceSheet
// 	if err := u.db.
// 		Where("month = ?", time.Now().Month()).
// 		Where("account_name = ?", "Pembelian").
// 		First(&pembelian).Error; err != nil {
// 		return err
// 	}

// 	var persediaan models.BalanceSheet
// 	if err := u.db.
// 		Where("month = ?", time.Now().Month()).
// 		Where("account_name = ?", "Persediaan Barang Dagang").
// 		First(&persediaan).Error; err != nil {
// 		return err
// 	}

// 	var hutang models.BalanceSheet
// 	if err := u.db.
// 		Where("month = ?", time.Now().Month()).
// 		Where("account_name = ?", "Hutang Dagang").
// 		First(&hutang).Error; err != nil {
// 		return err
// 	}

// 	totalPembelianCash := 0.0
// 	totalPembelianNonCash := 0.0

// 	for _, value := range transactions {

// 		if value.Payment == "Tunai" {
// 			totalPembelianCash += float64(value.Total)
// 		}

// 		if value.Payment == "Non-Tunai" {
// 			totalPembelianNonCash += float64(value.Total)
// 		}
// 	}

// 	totalPembelian := totalPembelianCash + totalPembelianNonCash

// 	var kasBalanceInDb models.Balance
// 	var pembelianBalanceInDb models.Balance
// 	var hutangBalanceInDb models.Balance
// 	var persediaanBalanceInDb models.Balance

// 	err := json.Unmarshal(kas.Balance, &kasBalanceInDb)

// 	if err != nil {
// 		return err
// 	}

// 	err = json.Unmarshal(pembelian.Balance, &pembelianBalanceInDb)

// 	if err != nil {
// 		return err
// 	}

// 	err = json.Unmarshal(hutang.Balance, &hutangBalanceInDb)

// 	if err != nil {
// 		return err
// 	}

// 	err = json.Unmarshal(persediaan.Balance, &persediaanBalanceInDb)

// 	if err != nil {
// 		return err
// 	}

// 	newKas := models.Balance{
// 		Debit:  kasBalanceInDb.Debit - totalPembelianCash,
// 		Credit: kasBalanceInDb.Credit,
// 	}

// 	newPembelian := models.Balance{
// 		Debit:  pembelianBalanceInDb.Debit + totalPembelian,
// 		Credit: pembelianBalanceInDb.Credit,
// 	}

// 	newHutang := models.Balance{
// 		Debit:  hutangBalanceInDb.Debit,
// 		Credit: hutangBalanceInDb.Credit + totalPembelianNonCash,
// 	}

// 	newPersediaan := models.Balance{
// 		Debit:  persediaanBalanceInDb.Debit + totalPembelian,
// 		Credit: persediaanBalanceInDb.Credit,
// 	}

// 	kasFormatted, err := json.Marshal(&newKas)

// 	if err != nil {
// 		return err
// 	}

// 	pembelianFormatted, err := json.Marshal(&newPembelian)

// 	if err != nil {
// 		return err
// 	}

// 	hutangFormatted, err := json.Marshal(&newHutang)

// 	if err != nil {
// 		return err
// 	}

// 	persediaanFormatted, err := json.Marshal(&newPersediaan)

// 	if err != nil {
// 		return err
// 	}

// 	kas.Balance = kasFormatted
// 	pembelian.Balance = pembelianFormatted
// 	hutang.Balance = hutangFormatted
// 	persediaan.Balance = persediaanFormatted

// 	if err := u.db.Save(&kas).Error; err != nil {
// 		return err
// 	}

// 	if err := u.db.Save(&pembelian).Error; err != nil {
// 		return err
// 	}

// 	if err := u.db.Save(&hutang).Error; err != nil {
// 		return err
// 	}

// 	if err := u.db.Save(&persediaan).Error; err != nil {
// 		return err
// 	}

// 	for i := range transactions {
// 		transactions[i].AlreadyPosted = 1
// 		if err := u.db.Model(&transactions[i]).Update("already_posted", 1).Error; err != nil {
// 			log.Println("Error Updating Transactions:", err)
// 			continue
// 		}
// 	}

// 	fmt.Println("Posting Completed!")
// 	return nil
// }
