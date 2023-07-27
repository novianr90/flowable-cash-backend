package readSale

import (
	"flowable-cash-backend/models"

	"gorm.io/gorm"
)

type Repository interface {
	// Read: Tx Sale Type
	ReadAllSaleTypeTransactions() (*[]models.Pemasukkan, error)
	ReadSaleTypeById(input *models.Pemasukkan) (*models.Pemasukkan, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryRead(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) ReadAllSaleTypeTransactions() (*[]models.Pemasukkan, error) {

	var saleTransactions []models.Pemasukkan

	db := r.db.Model(&models.Pemasukkan{})

	err := db.Find(&saleTransactions).Error

	if err != nil {
		return &[]models.Pemasukkan{}, err
	}

	return &saleTransactions, nil
}

func (r *repository) ReadSaleTypeById(input *models.Pemasukkan) (*models.Pemasukkan, error) {

	var saleTransaction models.Pemasukkan

	db := r.db.Model(&models.Pemasukkan{})

	err := db.Where("id = ?", input.ID).First(&saleTransaction).Error

	if err != nil {
		return &models.Pemasukkan{}, err
	}

	return &saleTransaction, nil
}
