package readTransaction

import (
	"flowable-cash-backend/models"

	"gorm.io/gorm"
)

type Repository interface {
	ReadAllTransactions() (*[]models.Pemasukkan, *[]models.Pengeluaran, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryRead(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) ReadAllTransactions() (*[]models.Pemasukkan, *[]models.Pengeluaran, error) {
	var pengeluaran []models.Pengeluaran
	var pemasukkan []models.Pemasukkan

	err := r.db.Find(&pengeluaran).Error

	if err != nil {
		return &[]models.Pemasukkan{}, &[]models.Pengeluaran{}, err
	}

	err = r.db.Find(&pemasukkan).Error

	if err != nil {
		return &[]models.Pemasukkan{}, &[]models.Pengeluaran{}, err
	}

	return &pemasukkan, &pengeluaran, nil
}
