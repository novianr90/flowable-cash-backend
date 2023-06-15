package updateTransaction

import (
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

	db := r.db.Model(models.Transaction{})

	newTransaction := models.Transaction{
		Name:        input.Name,
		Date:        input.Date,
		Type:        input.Type,
		Total:       input.Total,
		FeeType:     input.FeeType,
		Fee:         input.Fee,
		Description: input.Description,
	}

	updateTransaction := db.Where("id = ?", input.ID).Updates(&newTransaction)

	if updateTransaction.RowsAffected == 0 {
		return &models.Transaction{}, updateTransaction.Error
	}

	if updateTransaction.Error != nil {
		return &models.Transaction{}, updateTransaction.Error
	}

	_ = db.Where("id = ?", input.ID).First(&transaction)

	return &transaction, nil
}
