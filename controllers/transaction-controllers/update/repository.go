package updateTransaction

import (
	"errors"
	"flowable-cash-backend/models"

	"gorm.io/gorm"
)

type Repository interface {
	UpdateTransactionRepository(input *models.Transaction) (*models.Transaction, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryUpdate(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) UpdateTransactionRepository(input *models.Transaction) (*models.Transaction, error) {
	var transaction models.Transaction

	db := r.db.Model(&transaction)

	checkTransactionExist := db.Where("id = ?", input.ID).Find(&transaction)

	if checkTransactionExist.RowsAffected < 1 {
		return &transaction, errors.New("no data found")
	}

	transaction = models.Transaction{
		Name:        transaction.Name,
		Date:        transaction.Date,
		Type:        transaction.Type,
		Total:       transaction.Total,
		Description: transaction.Description,
	}

	var response models.Transaction

	updateTransaction := db.Where("id = ?", input.ID).Updates(&response)

	if updateTransaction.Error != nil {
		return &models.Transaction{}, updateTransaction.Error
	}

	return &response, nil
}
