package readSale

import (
	"flowable-cash-backend/models"

	"gorm.io/gorm"
)

type Repository interface {
	// Read: Tx Sale Type
	ReadAllSaleTypeTransactions() (*[]models.Transaction, error)
	ReadSaleTypeById(input *models.Transaction) (*models.Transaction, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryRead(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) ReadAllSaleTypeTransactions() (*[]models.Transaction, error) {
	typeSale := "Sale"

	var saleTransactions []models.Transaction

	db := r.db.Model(&models.Transaction{})

	err := db.Where("type = ?", typeSale).Find(&saleTransactions).Error

	if err != nil {
		return &[]models.Transaction{}, err
	}

	return &saleTransactions, nil
}

func (r *repository) ReadSaleTypeById(input *models.Transaction) (*models.Transaction, error) {
	typeSale := "Sale"

	var saleTransaction models.Transaction

	db := r.db.Model(&models.Transaction{})

	err := db.Where("type = ?", typeSale).Where("id = ?", input.ID).First(&saleTransaction).Error

	if err != nil {
		return &models.Transaction{}, err
	}

	return &saleTransaction, nil
}
