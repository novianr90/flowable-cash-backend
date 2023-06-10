package createTransaction

import (
	"flowable-cash-backend/models"

	"gorm.io/gorm"
)

type Repository interface {
	CreateTransactionRepository(input *models.Transaction) (*models.Transaction, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryCreate(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) CreateTransactionRepository(input *models.Transaction) (*models.Transaction, error) {

	transaction := models.Transaction{
		Name:        input.Name,
		Date:        input.Date,
		Type:        input.Type,
		Total:       input.Total,
		Description: input.Description,
		FeeType:     input.FeeType,
		Fee:         input.Fee,
	}

	if err := r.db.Create(&transaction).Error; err != nil {
		return &models.Transaction{}, err
	}

	return &transaction, nil
}
