package updateTransaction

import (
	"errors"
	"flowable-cash-backend/models"

	"gorm.io/gorm"
)

type Repository interface {
	UpdateTransactionRepository(trxType string, query interface{}) error
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryUpdate(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) UpdateTransactionRepository(trxType string, query interface{}) error {

	if query == nil {
		return errors.New("there's no data to update")
	}

	if trxType == "Pemasukkan" {

		updatePemasukkan := query.(models.Pemasukkan)

		res := r.db.Model(&models.Pemasukkan{}).Where("id = ?", updatePemasukkan.ID).Updates(updatePemasukkan)

		if res.Error != nil {
			return res.Error
		}

		if res.RowsAffected == 0 {
			return errors.New("error when updating")
		}

		return nil

	} else {
		updatePengeluaran := query.(models.Pengeluaran)

		res := r.db.Model(&models.Pengeluaran{}).Where("id = ?", updatePengeluaran.ID).Updates(updatePengeluaran)

		if res.Error != nil {
			return res.Error
		}

		if res.RowsAffected == 0 {
			return errors.New("error when updating")
		}

		return nil
	}
}
