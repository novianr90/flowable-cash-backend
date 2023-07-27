package deleteTransaction

import (
	"errors"
	"flowable-cash-backend/models"

	"gorm.io/gorm"
)

type Repository interface {
	DeletePemasukkan(input *models.Pemasukkan) error
	DeletePengeluaran(input *models.Pengeluaran) error
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryDelete(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) DeletePemasukkan(input *models.Pemasukkan) error {

	deleteTransactionId := r.db.Model(&models.Pemasukkan{}).Where("id = ?", input).Delete(&models.Pemasukkan{})

	if deleteTransactionId.Error != nil {
		return deleteTransactionId.Error
	}

	if deleteTransactionId.RowsAffected == 0 {
		return errors.New("no data to delete")
	}

	return nil
}

func (r *repository) DeletePengeluaran(input *models.Pengeluaran) error {

	deleteTransactionId := r.db.Model(&models.Pengeluaran{}).Where("id = ?", input).Delete(&models.Pemasukkan{})

	if deleteTransactionId.Error != nil {
		return deleteTransactionId.Error
	}

	if deleteTransactionId.RowsAffected == 0 {
		return errors.New("no data to delete")
	}

	return nil
}
