package deleteTransaction

import (
	"errors"
	"flowable-cash-backend/models"

	"gorm.io/gorm"
)

type Repository interface {
	DeleteTransactionRepository(input uint, trxType string) error
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryDelete(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) DeleteTransactionRepository(input uint, trxType string) error {

	if trxType == "" {
		return errors.New("please specify the type")
	}

	if trxType == "Pemasukkan" {

		var container models.Pemasukkan

		checkTransaction := r.db.Model(&models.Pemasukkan{}).Select("*").Where("id = ?", input).Find(&container)

		if checkTransaction.RowsAffected < 1 {
			return errors.New("no data found")
		}

		deleteTransactionId := r.db.Model(&models.Pemasukkan{}).Select("*").Where("id = ?", input).Find(&container).Delete(&container)

		if deleteTransactionId.Error != nil {
			return deleteTransactionId.Error
		}

		return nil
	} else {
		var container models.Pengeluaran

		checkTransaction := r.db.Model(&models.Pengeluaran{}).Select("*").Where("id = ?", input).Find(&container)

		if checkTransaction.RowsAffected < 1 {
			return errors.New("no data found")
		}

		deleteTransactionId := r.db.Model(&models.Pengeluaran{}).Select("*").Where("id = ?", input).Find(&container).Delete(&container)

		if deleteTransactionId.Error != nil {
			return deleteTransactionId.Error
		}

		return nil
	}
}
