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
	PostingPemasukkan() error
	PostingBahanBaku() error
	PostingBarangDagang() error
	PostingBahanTambahan() error
	PostingPeralatan() error
	PostingBayarHutang() error
	PostingBayarPiutang() error
	PostingBiayaBiaya() error
}

type usecase struct {
	db *gorm.DB
}

func NewUseCaseService(db *gorm.DB) *usecase {
	return &usecase{db: db}
}

func (u *usecase) PostingPemasukkan() error {

	var transactions []models.Pemasukkan

	if err := u.db.
		Model(&models.Pemasukkan{}).
		Where("already_posted = ?", 0).
		Where("name = ?", "Pendapatan").
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
		if err := u.db.Model(&transactions[i]).Update("already_posted", transactions[i].AlreadyPosted).Error; err != nil {
			log.Println("Error Updating Transactions:", err)
			continue
		}
	}

	fmt.Println("Posting Completed!")
	return nil
}

func (u *usecase) PostingBahanBaku() error {
	var pengeluarans []models.Pengeluaran

	if err := u.db.Model(&models.Pengeluaran{}).
		Where("already_posted = ?", 0).
		Where("name = ?", "Bahan Baku").
		Find(&pengeluarans).Error; err != nil {
		return err
	}

	if pengeluarans == nil {
		return errors.New("no remaining data to posting")
	}

	var kas models.BalanceSheet
	if err := u.db.Model(&models.BalanceSheet{}).
		Where("month = ?", time.Now().Month()).
		Where("account_name = ?", "Kas").
		First(&kas).Error; err != nil {
		return err
	}

	var hutang models.BalanceSheet
	if err := u.db.Model(&models.BalanceSheet{}).
		Where("month = ?", time.Now().Month()).
		Where("account_name = ?", "Hutang Dagang").
		First(&hutang).Error; err != nil {
		return err
	}

	var bahanBaku models.BalanceSheet
	if err := u.db.Model(&models.BalanceSheet{}).
		Where("month = ?", time.Now().Month()).
		Where("account_name = ?", "Bahan Baku").
		First(&bahanBaku).Error; err != nil {
		return err
	}

	var modal models.BalanceSheet
	if err := u.db.Model(&models.BalanceSheet{}).
		Where("month = ?", time.Now().Month()).
		Where("account_name = ?", "Modal").
		First(&modal).Error; err != nil {
		return err
	}

	var kasBalance models.Balance
	var hutangBalance models.Balance
	var bahanBakuBalance models.Balance
	var modalBalance models.Balance

	_ = json.Unmarshal(kas.Balance, &kasBalance)
	_ = json.Unmarshal(hutang.Balance, &hutangBalance)
	_ = json.Unmarshal(bahanBaku.Balance, &bahanBakuBalance)
	_ = json.Unmarshal(modal.Balance, &modalBalance)

	totalCash := 0.0
	totalNonCash := 0.0

	for _, value := range pengeluarans {
		if value.Payment == "Tunai" {
			totalCash += float64(value.Total)
		} else {
			totalNonCash += float64(value.Total)
		}
	}

	totalKas := kasBalance.Debit - totalCash
	totalBahanBaku := totalCash + totalNonCash

	totalCashBalanceRaw := models.Balance{
		Debit:  totalKas,
		Credit: 0,
	}

	totalNonCashBalanceRaw := models.Balance{
		Debit:  0,
		Credit: hutangBalance.Credit + totalNonCash,
	}

	totalRaw := models.Balance{
		Debit:  bahanBakuBalance.Debit + totalBahanBaku,
		Credit: 0,
	}

	totalModal := models.Balance{
		Debit:  0,
		Credit: modalBalance.Credit - totalCash,
	}

	cashFormatted, _ := json.Marshal(&totalCashBalanceRaw)
	nonCashFormatted, _ := json.Marshal(&totalNonCashBalanceRaw)
	totalFormatted, _ := json.Marshal(&totalRaw)
	modalFormatted, _ := json.Marshal(&totalModal)

	kas.Balance = cashFormatted
	hutang.Balance = nonCashFormatted
	bahanBaku.Balance = totalFormatted
	modal.Balance = modalFormatted

	if err := u.db.Model(&models.BalanceSheet{}).
		Where("month = ?", time.Now().Month()).
		Where("account_name = ?", "Kas").
		Save(&kas).Error; err != nil {
		return err
	}

	if err := u.db.Model(&models.BalanceSheet{}).
		Where("month = ?", time.Now().Month()).
		Where("account_name = ?", "Hutang Dagang").
		Save(&hutang).Error; err != nil {
		return err
	}

	if err := u.db.Model(&models.BalanceSheet{}).
		Where("month = ?", time.Now().Month()).
		Where("account_name = ?", "Bahan Baku").
		Save(&bahanBaku).Error; err != nil {
		return err
	}

	if err := u.db.Model(&models.BalanceSheet{}).
		Where("month = ?", time.Now().Month()).
		Where("account_name = ?", "Modal").
		Save(&modal).Error; err != nil {
		return err
	}

	for i := range pengeluarans {
		pengeluarans[i].AlreadyPosted = 1
		if err := u.db.Model(&pengeluarans[i]).Update("already_posted", pengeluarans[i].AlreadyPosted).Error; err != nil {
			log.Println("Error Updating Transactions:", err)
			continue
		}
	}

	// old logic
	// if err := u.db.Model(&models.Pengeluaran{}).
	// 	Where("name = ?", "Bahan Baku").
	// 	Save(&pengeluarans).Error; err != nil {
	// 	return err
	// }

	log.Println("Posting Bahan Baku Complete!")

	return nil
}

func (u *usecase) PostingBarangDagang() error {
	var pengeluarans []models.Pengeluaran

	if err := u.db.Model(&models.Pengeluaran{}).
		Where("already_posted = ?", 0).
		Where("name = ?", "Barang Dagang").
		Find(&pengeluarans).Error; err != nil {
		return err
	}

	if pengeluarans == nil {
		return errors.New("no remaining data to posting")
	}

	var kas models.BalanceSheet
	if err := u.db.Model(&models.BalanceSheet{}).
		Where("month = ?", time.Now().Month()).
		Where("account_name = ?", "Kas").
		First(&kas).Error; err != nil {
		return err
	}

	var hutang models.BalanceSheet
	if err := u.db.Model(&models.BalanceSheet{}).
		Where("month = ?", time.Now().Month()).
		Where("account_name = ?", "Hutang Dagang").
		First(&hutang).Error; err != nil {
		return err
	}

	var barangDagang models.BalanceSheet
	if err := u.db.Model(&models.BalanceSheet{}).
		Where("month = ?", time.Now().Month()).
		Where("account_name = ?", "Barang Dagang").
		First(&barangDagang).Error; err != nil {
		return err
	}

	var modal models.BalanceSheet
	if err := u.db.Model(&models.BalanceSheet{}).
		Where("month = ?", time.Now().Month()).
		Where("account_name = ?", "Modal").
		First(&modal).Error; err != nil {
		return err
	}

	var kasBalance models.Balance
	var hutangBalance models.Balance
	var barangDagangBalance models.Balance
	var modalBalance models.Balance

	_ = json.Unmarshal(kas.Balance, &kasBalance)
	_ = json.Unmarshal(hutang.Balance, &hutangBalance)
	_ = json.Unmarshal(barangDagang.Balance, &barangDagangBalance)
	_ = json.Unmarshal(modal.Balance, &modalBalance)

	totalCash := 0.0
	totalNonCash := 0.0

	for _, value := range pengeluarans {
		if value.Payment == "Tunai" {
			totalCash += float64(value.Total)
		} else {
			totalNonCash += float64(value.Total)
		}
	}

	totalBarangDagang := totalCash + totalNonCash

	totalCashBalanceRaw := models.Balance{
		Debit:  kasBalance.Debit - totalCash,
		Credit: 0,
	}

	totalNonCashBalanceRaw := models.Balance{
		Debit:  0,
		Credit: hutangBalance.Credit + totalNonCash,
	}

	totalRaw := models.Balance{
		Debit:  barangDagangBalance.Debit + totalBarangDagang,
		Credit: 0,
	}

	totalModal := models.Balance{
		Debit:  0,
		Credit: modalBalance.Credit - totalCash,
	}

	cashFormatted, _ := json.Marshal(&totalCashBalanceRaw)
	nonCashFormatted, _ := json.Marshal(&totalNonCashBalanceRaw)
	totalFormatted, _ := json.Marshal(&totalRaw)
	modalFormatted, _ := json.Marshal(&totalModal)

	kas.Balance = cashFormatted
	hutang.Balance = nonCashFormatted
	barangDagang.Balance = totalFormatted
	modal.Balance = modalFormatted

	// for _, value := range pengeluarans {
	// 	value.AlreadyPosted = 1
	// }

	if err := u.db.Model(&models.BalanceSheet{}).
		Where("month = ?", time.Now().Month()).
		Where("account_name = ?", "Kas").
		Save(&kas).Error; err != nil {
		return err
	}

	if err := u.db.Model(&models.BalanceSheet{}).
		Where("month = ?", time.Now().Month()).
		Where("account_name = ?", "Hutang Dagang").
		Save(&hutang).Error; err != nil {
		return err
	}

	if err := u.db.Model(&models.BalanceSheet{}).
		Where("month = ?", time.Now().Month()).
		Where("account_name = ?", "Barang Dagang").
		Save(&barangDagang).Error; err != nil {
		return err
	}

	if err := u.db.Model(&models.BalanceSheet{}).
		Where("month = ?", time.Now().Month()).
		Where("account_name = ?", "Modal").
		Save(&modal).Error; err != nil {
		return err
	}

	for i := range pengeluarans {
		pengeluarans[i].AlreadyPosted = 1
		if err := u.db.Model(&pengeluarans[i]).Update("already_posted", pengeluarans[i].AlreadyPosted).Error; err != nil {
			log.Println("Error Updating Transactions:", err)
			continue
		}
	}

	log.Println("Posting Barang Dagang Complete!")

	return nil
}

func (u *usecase) PostingBahanTambahan() error {
	var pengeluarans []models.Pengeluaran

	if err := u.db.Model(&models.Pengeluaran{}).
		Where("already_posted = ?", 0).
		Where("name = ?", "Bahan Tambahan").
		Find(&pengeluarans).Error; err != nil {
		return err
	}

	if pengeluarans == nil {
		return errors.New("no remaining data to posting")
	}

	var kas models.BalanceSheet
	if err := u.db.Model(&models.BalanceSheet{}).
		Where("month = ?", time.Now().Month()).
		Where("account_name = ?", "Kas").
		First(&kas).Error; err != nil {
		return err
	}

	var hutang models.BalanceSheet
	if err := u.db.Model(&models.BalanceSheet{}).
		Where("month = ?", time.Now().Month()).
		Where("account_name = ?", "Hutang Dagang").
		First(&hutang).Error; err != nil {
		return err
	}

	var bahanTambahan models.BalanceSheet
	if err := u.db.Model(&models.BalanceSheet{}).
		Where("month = ?", time.Now().Month()).
		Where("account_name = ?", "Bahan Tambahan").
		First(&bahanTambahan).Error; err != nil {
		return err
	}

	var modal models.BalanceSheet
	if err := u.db.Model(&models.BalanceSheet{}).
		Where("month = ?", time.Now().Month()).
		Where("account_name = ?", "Modal").
		First(&modal).Error; err != nil {
		return err
	}

	var kasBalance models.Balance
	var hutangBalance models.Balance
	var bahanTambahanBalance models.Balance
	var modalBalance models.Balance

	_ = json.Unmarshal(kas.Balance, &kasBalance)
	_ = json.Unmarshal(hutang.Balance, &hutangBalance)
	_ = json.Unmarshal(bahanTambahan.Balance, &bahanTambahanBalance)
	_ = json.Unmarshal(modal.Balance, &modalBalance)

	totalCash := 0.0
	totalNonCash := 0.0

	for _, value := range pengeluarans {
		if value.Payment == "Tunai" {
			totalCash += float64(value.Total)
		} else {
			totalNonCash += float64(value.Total)
		}
	}

	totalBahanTambahan := totalCash + totalNonCash

	totalCashBalanceRaw := models.Balance{
		Debit:  kasBalance.Debit - totalCash,
		Credit: 0,
	}

	totalNonCashBalanceRaw := models.Balance{
		Debit:  0,
		Credit: hutangBalance.Credit + totalNonCash,
	}

	totalRaw := models.Balance{
		Debit:  bahanTambahanBalance.Debit + totalBahanTambahan,
		Credit: 0,
	}

	totalModal := models.Balance{
		Debit:  0,
		Credit: modalBalance.Credit - totalCash,
	}

	cashFormatted, _ := json.Marshal(&totalCashBalanceRaw)
	nonCashFormatted, _ := json.Marshal(&totalNonCashBalanceRaw)
	totalFormatted, _ := json.Marshal(&totalRaw)
	modalFormatted, _ := json.Marshal(&totalModal)

	kas.Balance = cashFormatted
	hutang.Balance = nonCashFormatted
	bahanTambahan.Balance = totalFormatted
	modal.Balance = modalFormatted

	// for _, value := range pengeluarans {
	// 	value.AlreadyPosted = 1
	// }

	if err := u.db.Model(&models.BalanceSheet{}).
		Where("month = ?", time.Now().Month()).
		Where("account_name = ?", "Kas").
		Save(&kas).Error; err != nil {
		return err
	}

	if err := u.db.Model(&models.BalanceSheet{}).
		Where("month = ?", time.Now().Month()).
		Where("account_name = ?", "Hutang Dagang").
		Save(&hutang).Error; err != nil {
		return err
	}

	if err := u.db.Model(&models.BalanceSheet{}).
		Where("month = ?", time.Now().Month()).
		Where("account_name = ?", "Bahan Tambahan").
		Save(&bahanTambahan).Error; err != nil {
		return err
	}

	if err := u.db.Model(&models.BalanceSheet{}).
		Where("month = ?", time.Now().Month()).
		Where("account_name = ?", "Modal").
		Save(&modal).Error; err != nil {
		return err
	}

	for i := range pengeluarans {
		pengeluarans[i].AlreadyPosted = 1
		if err := u.db.Model(&pengeluarans[i]).Update("already_posted", pengeluarans[i].AlreadyPosted).Error; err != nil {
			log.Println("Error Updating Transactions:", err)
			continue
		}
	}

	// old
	// if err := u.db.Model(&models.Pengeluaran{}).
	// 	Where("name = ?", "Bahan Tambahan").
	// 	Save(&pengeluarans).Error; err != nil {
	// 	return err
	// }

	log.Println("Posting Bahan Tambahan Complete!")

	return nil
}

func (u *usecase) PostingPeralatan() error {
	var pengeluarans []models.Pengeluaran

	if err := u.db.Model(&models.Pengeluaran{}).
		Where("already_posted = ?", 0).
		Where("name = ?", "Pengeluaran untuk Pembelian Alat Usaha").
		Find(&pengeluarans).Error; err != nil {
		return err
	}

	if pengeluarans == nil {
		return errors.New("no remaining data to posting")
	}

	var kas models.BalanceSheet
	if err := u.db.Model(&models.BalanceSheet{}).
		Where("month = ?", time.Now().Month()).
		Where("account_name = ?", "Kas").
		First(&kas).Error; err != nil {
		return err
	}

	var hutang models.BalanceSheet
	if err := u.db.Model(&models.BalanceSheet{}).
		Where("month = ?", time.Now().Month()).
		Where("account_name = ?", "Hutang Dagang").
		First(&hutang).Error; err != nil {
		return err
	}

	var peralatan models.BalanceSheet
	if err := u.db.Model(&models.BalanceSheet{}).
		Where("month = ?", time.Now().Month()).
		Where("account_name = ?", "Peralatan").
		First(&peralatan).Error; err != nil {
		return err
	}

	var modal models.BalanceSheet
	if err := u.db.Model(&models.BalanceSheet{}).
		Where("month = ?", time.Now().Month()).
		Where("account_name = ?", "Modal").
		First(&modal).Error; err != nil {
		return err
	}

	var kasBalance models.Balance
	var hutangBalance models.Balance
	var peralatanBalance models.Balance
	var modalBalance models.Balance

	_ = json.Unmarshal(kas.Balance, &kasBalance)
	_ = json.Unmarshal(hutang.Balance, &hutangBalance)
	_ = json.Unmarshal(peralatan.Balance, &peralatanBalance)
	_ = json.Unmarshal(modal.Balance, &modalBalance)

	totalCash := 0.0
	totalNonCash := 0.0

	for _, value := range pengeluarans {
		if value.Payment == "Tunai" {
			totalCash += float64(value.Total)
		} else {
			totalNonCash += float64(value.Total)
		}
	}

	totalPeralatan := totalCash + totalNonCash

	totalCashBalanceRaw := models.Balance{
		Debit:  kasBalance.Debit - totalCash,
		Credit: 0,
	}

	totalNonCashBalanceRaw := models.Balance{
		Debit:  0,
		Credit: hutangBalance.Credit + totalNonCash,
	}

	totalRaw := models.Balance{
		Debit:  peralatanBalance.Debit + totalPeralatan,
		Credit: 0,
	}

	totalModal := models.Balance{
		Debit:  0,
		Credit: modalBalance.Credit - totalCash,
	}

	cashFormatted, _ := json.Marshal(&totalCashBalanceRaw)
	nonCashFormatted, _ := json.Marshal(&totalNonCashBalanceRaw)
	totalFormatted, _ := json.Marshal(&totalRaw)
	modalFormatted, _ := json.Marshal(&totalModal)

	kas.Balance = cashFormatted
	hutang.Balance = nonCashFormatted
	peralatan.Balance = totalFormatted
	modal.Balance = modalFormatted

	// for _, value := range pengeluarans {
	// 	value.AlreadyPosted = 1
	// }

	if err := u.db.Model(&models.BalanceSheet{}).
		Where("month = ?", time.Now().Month()).
		Where("account_name = ?", "Kas").
		Save(&kas).Error; err != nil {
		return err
	}

	if err := u.db.Model(&models.BalanceSheet{}).
		Where("month = ?", time.Now().Month()).
		Where("account_name = ?", "Hutang Dagang").
		Save(&hutang).Error; err != nil {
		return err
	}

	if err := u.db.Model(&models.BalanceSheet{}).
		Where("month = ?", time.Now().Month()).
		Where("account_name = ?", "Peralatan").
		Save(&peralatan).Error; err != nil {
		return err
	}

	if err := u.db.Model(&models.BalanceSheet{}).
		Where("month = ?", time.Now().Month()).
		Where("account_name = ?", "Modal").
		Save(&modal).Error; err != nil {
		return err
	}

	for i := range pengeluarans {
		pengeluarans[i].AlreadyPosted = 1
		if err := u.db.Model(&pengeluarans[i]).Update("already_posted", pengeluarans[i].AlreadyPosted).Error; err != nil {
			log.Println("Error Updating Transactions:", err)
			continue
		}
	}

	// OLD
	// if err := u.db.Model(&models.Pengeluaran{}).
	// 	Where("name = ?", "Pengeluaran untuk Pembelian Alat Usaha").
	// 	Save(&pengeluarans).Error; err != nil {
	// 	return err
	// }

	log.Println("Posting Peralatan Complete!")

	return nil
}

func (u *usecase) PostingBayarHutang() error {
	var pengeluarans []models.Pengeluaran

	if err := u.db.Model(&models.Pengeluaran{}).
		Where("already_posted = ?", 0).
		Where("name = ?", "Membayar Hutang").
		Find(&pengeluarans).Error; err != nil {
		return err
	}

	if pengeluarans == nil {
		return errors.New("no remaining data to posting")
	}

	var kas models.BalanceSheet
	if err := u.db.Model(&models.BalanceSheet{}).
		Where("month = ?", time.Now().Month()).
		Where("account_name = ?", "Kas").
		First(&kas).Error; err != nil {
		return err
	}

	var hutang models.BalanceSheet
	if err := u.db.Model(&models.BalanceSheet{}).
		Where("month = ?", time.Now().Month()).
		Where("account_name = ?", "Hutang Dagang").
		First(&hutang).Error; err != nil {
		return err
	}

	var modal models.BalanceSheet
	if err := u.db.Model(&models.BalanceSheet{}).
		Where("month = ?", time.Now().Month()).
		Where("account_name = ?", "Modal").
		First(&modal).Error; err != nil {
		return err
	}

	var kasBalance models.Balance
	var hutangBalance models.Balance
	var modalBalance models.Balance

	_ = json.Unmarshal(kas.Balance, &kasBalance)
	_ = json.Unmarshal(hutang.Balance, &hutangBalance)
	_ = json.Unmarshal(modal.Balance, &modalBalance)

	total := 0.0

	for _, value := range pengeluarans {
		if value.Payment == "Tunai" {
			total += float64(value.Total)
		}
	}

	totalCashBalanceRaw := models.Balance{
		Debit:  kasBalance.Debit - total,
		Credit: 0,
	}

	totalNonCashBalanceRaw := models.Balance{
		Debit:  0,
		Credit: hutangBalance.Credit - total,
	}

	totalModal := models.Balance{
		Debit:  0,
		Credit: modalBalance.Credit - total,
	}

	cashFormatted, _ := json.Marshal(&totalCashBalanceRaw)
	nonCashFormatted, _ := json.Marshal(&totalNonCashBalanceRaw)
	modalFormatted, _ := json.Marshal(&totalModal)

	kas.Balance = cashFormatted
	hutang.Balance = nonCashFormatted
	modal.Balance = modalFormatted

	for _, value := range pengeluarans {
		value.AlreadyPosted = 1
	}

	if err := u.db.Model(&models.BalanceSheet{}).
		Where("month = ?", time.Now().Month()).
		Where("account_name = ?", "Kas").
		Save(&kas).Error; err != nil {
		return err
	}

	if err := u.db.Model(&models.BalanceSheet{}).
		Where("month = ?", time.Now().Month()).
		Where("account_name = ?", "Hutang Dagang").
		Save(&hutang).Error; err != nil {
		return err
	}

	if err := u.db.Model(&models.BalanceSheet{}).
		Where("month = ?", time.Now().Month()).
		Where("account_name = ?", "Modal").
		Save(&modal).Error; err != nil {
		return err
	}

	for i := range pengeluarans {
		pengeluarans[i].AlreadyPosted = 1
		if err := u.db.Model(&pengeluarans[i]).Update("already_posted", pengeluarans[i].AlreadyPosted).Error; err != nil {
			log.Println("Error Updating Transactions:", err)
			continue
		}
	}

	// if err := u.db.Model(&models.Pengeluaran{}).
	// 	Where("name = ?", "Membayar Hutang").
	// 	Save(&pengeluarans).Error; err != nil {
	// 	return err
	// }

	log.Println("Posting Hutang Complete!")

	return nil
}

func (u *usecase) PostingBayarPiutang() error {
	var pemasukkan []models.Pemasukkan

	if err := u.db.Model(&models.Pemasukkan{}).
		Where("already_posted = ?", 0).
		Where("name = ?", "Pembayaran Piutang").
		Find(&pemasukkan).Error; err != nil {
		return err
	}

	if pemasukkan == nil {
		return errors.New("no remaining data to posting")
	}

	var kas models.BalanceSheet
	if err := u.db.Model(&models.BalanceSheet{}).
		Where("month = ?", time.Now().Month()).
		Where("account_name = ?", "Kas").
		First(&kas).Error; err != nil {
		return err
	}

	var piutang models.BalanceSheet
	if err := u.db.Model(&models.BalanceSheet{}).
		Where("month = ?", time.Now().Month()).
		Where("account_name = ?", "Piutang Dagang").
		First(&piutang).Error; err != nil {
		return err
	}

	var modal models.BalanceSheet
	if err := u.db.Model(&models.BalanceSheet{}).
		Where("month = ?", time.Now().Month()).
		Where("account_name = ?", "Modal").
		First(&modal).Error; err != nil {
		return err
	}

	var kasBalance models.Balance
	var piutangBalance models.Balance
	var modalBalance models.Balance

	_ = json.Unmarshal(kas.Balance, &kasBalance)
	_ = json.Unmarshal(piutang.Balance, &piutangBalance)
	_ = json.Unmarshal(modal.Balance, &modalBalance)

	total := 0.0

	for _, value := range pemasukkan {
		if value.Payment == "Tunai" {
			total += float64(value.Total)
		}
	}

	totalCashBalanceRaw := models.Balance{
		Debit:  kasBalance.Debit + total,
		Credit: 0,
	}

	totalNonCashBalanceRaw := models.Balance{
		Debit:  piutangBalance.Debit - total,
		Credit: 0,
	}

	totalModal := models.Balance{
		Debit:  0,
		Credit: modalBalance.Credit + total,
	}

	cashFormatted, _ := json.Marshal(&totalCashBalanceRaw)
	nonCashFormatted, _ := json.Marshal(&totalNonCashBalanceRaw)
	modalFormatted, _ := json.Marshal(&totalModal)

	kas.Balance = cashFormatted
	piutang.Balance = nonCashFormatted
	modal.Balance = modalFormatted

	// for _, value := range pemasukkan {
	// 	value.AlreadyPosted = 1
	// }

	if err := u.db.Model(&models.BalanceSheet{}).
		Where("month = ?", time.Now().Month()).
		Where("account_name = ?", "Kas").
		Save(&kas).Error; err != nil {
		return err
	}

	if err := u.db.Model(&models.BalanceSheet{}).
		Where("month = ?", time.Now().Month()).
		Where("account_name = ?", "Piutang Dagang").
		Save(&piutang).Error; err != nil {
		return err
	}

	if err := u.db.Model(&models.BalanceSheet{}).
		Where("month = ?", time.Now().Month()).
		Where("account_name = ?", "Modal").
		Save(&modal).Error; err != nil {
		return err
	}

	for i := range pemasukkan {
		pemasukkan[i].AlreadyPosted = 1
		if err := u.db.Model(&pemasukkan[i]).Update("already_posted", pemasukkan[i].AlreadyPosted).Error; err != nil {
			log.Println("Error Updating Transactions:", err)
			continue
		}
	}

	// OLD
	// if err := u.db.Model(&models.Pemasukkan{}).
	// 	Where("name = ?", "Pembayaran Piutang").
	// 	Save(&pemasukkan).Error; err != nil {
	// 	return err
	// }

	log.Println("Posting Piutang Complete!")

	return nil
}

func (u *usecase) PostingBiayaBiaya() error {
	var pengeluarans []models.Pengeluaran

	if err := u.db.Model(&models.Pengeluaran{}).
		Where("already_posted = ?", 0).
		Find(&pengeluarans).Error; err != nil {
		return err
	}

	if pengeluarans == nil {
		return errors.New("no remaining data to posting")
	}

	var kas models.BalanceSheet
	if err := u.db.Model(&models.BalanceSheet{}).
		Where("month = ?", time.Now().Month()).
		Where("account_name = ?", "Kas").
		First(&kas).Error; err != nil {
		return err
	}

	var hutang models.BalanceSheet
	if err := u.db.Model(&models.BalanceSheet{}).
		Where("month = ?", time.Now().Month()).
		Where("account_name = ?", "Hutang Dagang").
		First(&hutang).Error; err != nil {
		return err
	}

	var modal models.BalanceSheet
	if err := u.db.Model(&models.BalanceSheet{}).
		Where("month = ?", time.Now().Month()).
		Where("account_name = ?", "Modal").
		First(&modal).Error; err != nil {
		return err
	}

	var kasBalance models.Balance
	var hutangBalance models.Balance
	var modalBalance models.Balance

	_ = json.Unmarshal(kas.Balance, &kasBalance)
	_ = json.Unmarshal(hutang.Balance, &hutangBalance)
	_ = json.Unmarshal(modal.Balance, &modalBalance)

	totalCash := 0.0
	totalNonCash := 0.0

	filteredPengeluaran := filterBiaya(pengeluarans)

	for _, value := range filteredPengeluaran {
		if value.Payment == "Tunai" {
			totalCash += float64(value.Total)
		} else {
			totalNonCash += float64(value.Total)
		}
	}

	totalCashBalanceRaw := models.Balance{
		Debit:  kasBalance.Debit - totalCash,
		Credit: 0,
	}

	totalNonCashBalanceRaw := models.Balance{
		Debit:  0,
		Credit: hutangBalance.Credit + totalNonCash,
	}

	totalModal := models.Balance{
		Debit:  0,
		Credit: modalBalance.Credit - totalCash,
	}

	cashFormatted, _ := json.Marshal(&totalCashBalanceRaw)
	nonCashFormatted, _ := json.Marshal(&totalNonCashBalanceRaw)
	modalFormatted, _ := json.Marshal(&totalModal)

	kas.Balance = cashFormatted
	hutang.Balance = nonCashFormatted
	modal.Balance = modalFormatted

	// for _, value := range filteredPengeluaran {
	// 	value.AlreadyPosted = 1
	// }

	if err := u.db.Model(&models.BalanceSheet{}).
		Where("month = ?", time.Now().Month()).
		Where("account_name = ?", "Kas").
		Save(&kas).Error; err != nil {
		return err
	}

	if err := u.db.Model(&models.BalanceSheet{}).
		Where("month = ?", time.Now().Month()).
		Where("account_name = ?", "Hutang Dagang").
		Save(&hutang).Error; err != nil {
		return err
	}

	if err := u.db.Model(&models.BalanceSheet{}).
		Where("month = ?", time.Now().Month()).
		Where("account_name = ?", "Modal").
		Save(&modal).Error; err != nil {
		return err
	}

	for i := range filteredPengeluaran {
		filteredPengeluaran[i].AlreadyPosted = 1
		if err := u.db.Model(&filteredPengeluaran[i]).Update("already_posted", filteredPengeluaran[i].AlreadyPosted).Error; err != nil {
			log.Println("Error Updating Transactions:", err)
			continue
		}
	}

	// OLD
	// if err := u.db.Model(&models.Pengeluaran{}).
	// 	Save(&filteredPengeluaran).Error; err != nil {
	// 	return err
	// }

	log.Println("Posting Biaya Biaya Complete!")

	return nil
}

func filterBiaya(dataList []models.Pengeluaran) []models.Pengeluaran {
	var biayaList []models.Pengeluaran

	keyWordBiaya := map[string]struct{}{
		"Biaya Listrik":      {},
		"Biaya Air":          {},
		"Biaya Perbaikan":    {},
		"Biaya Promosi":      {},
		"Biaya Ongkos Kirim": {},
		"Biaya Pengemasan":   {},
		"Biaya Gaji":         {},
		"Biaya Sewa":         {},
		"Biaya Lainnya":      {},
	}

	for _, data := range dataList {
		_, isBiaya := keyWordBiaya[data.Name]
		if isBiaya {
			biayaList = append(biayaList, data)
		}
	}

	return biayaList
}
