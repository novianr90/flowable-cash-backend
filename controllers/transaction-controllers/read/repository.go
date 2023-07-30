package readTransaction

import (
	"flowable-cash-backend/models"

	"gorm.io/gorm"
)

type Repository interface {
	ReadAllTransactions() (*[]models.Pengeluaran, *[]models.Pemasukkan, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryRead(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) ReadAllTransactions() (*[]models.Pengeluaran, *[]models.Pemasukkan, error) {
	var pengeluaran []models.Pengeluaran
	var pemasukkan []models.Pemasukkan

	err := r.db.Find(&pengeluaran).Error

	if err != nil {
		return &[]models.Pengeluaran{}, &[]models.Pemasukkan{}, err
	}

	err = r.db.Find(&pemasukkan).Error

	if err != nil {
		return &[]models.Pengeluaran{}, &[]models.Pemasukkan{}, err
	}

	return &pengeluaran, &pemasukkan, nil
}
