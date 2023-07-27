package updateTransaction

import (
	"errors"
	"flowable-cash-backend/models"

	"gorm.io/gorm"
)

type Repository interface {
	UpdatePemasukkan(input *models.Pemasukkan) error
	UpdatePengeluaran(input *models.Pengeluaran) error
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryUpdate(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) UpdatePemasukkan(input *models.Pemasukkan) error {
	res := r.db.Model(&models.Pemasukkan{}).Where("id = ?", input.ID).Updates(&input)

	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return errors.New("error when updating")
	}

	return nil
}

func (r *repository) UpdatePengeluaran(input *models.Pengeluaran) error {
	res := r.db.Model(&models.Pengeluaran{}).Where("id = ?", input.ID).Updates(&input)

	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return errors.New("error when updating")
	}

	return nil
}
