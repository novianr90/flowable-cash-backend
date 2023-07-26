package readTransaction

import (
	"flowable-cash-backend/models"

	"gorm.io/gorm"
)

type Repository interface {
	ReadAllTransactions() (*[]models.Transaction, error)
	ReadTransactionById(input *models.Transaction) (*models.Transaction, error)
	ReadAllTransactionsByType(input *models.Transaction) (*[]models.Transaction, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryRead(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) ReadAllTransactions() (*[]models.Transaction, error) {
	var transactions []models.Transaction

	db := r.db.Model(&transactions)

	result := db.Find(&transactions)

	if result.Error != nil {
		return &transactions, result.Error
	}

	return &transactions, nil
}

func (r *repository) ReadTransactionById(input *models.Transaction) (*models.Transaction, error) {
	var transaction models.Transaction

	db := r.db.Model(&transaction)

	result := db.Where("id = ?", input.ID).First(&transaction)

	if result.Error != nil {
		return &transaction, result.Error
	}

	return &transaction, nil
}

func (r *repository) ReadAllTransactionsByType(input *models.Transaction) (*[]models.Transaction, error) {
	db := r.db.Model(&models.Transaction{})

	var transactions []models.Transaction

	err := db.Where("type = ?", input.Type).Find(&transactions).Error

	if err != nil {
		return &[]models.Transaction{}, err
	}

	return &transactions, nil
}
