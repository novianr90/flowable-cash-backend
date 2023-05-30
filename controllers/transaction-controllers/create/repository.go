package createTransaction

import (
	"flowable-cash-backend/internal/sorting"
	"flowable-cash-backend/models"

	"gorm.io/gorm"
)

type Repository interface {
	CreateTransactionRepository(input *models.Transaction) (*models.Transaction, error)
}

type repository struct {
	db      *gorm.DB
	sorting sorting.Sorting
}

func NewRepositoryCreate(db *gorm.DB, sorting sorting.Sorting) *repository {
	return &repository{db: db, sorting: sorting}
}

func (r *repository) CreateTransactionRepository(input *models.Transaction) (*models.Transaction, error) {

	transaction := models.Transaction{
		Name:        input.Name,
		Date:        input.Date,
		Type:        input.Type,
		Total:       input.Total,
		Description: input.Description,
	}

	if err := r.db.Create(&transaction).Error; err != nil {
		return &models.Transaction{}, err
	}

	if err := r.sorting.SortTransaction(); err != nil {
		return &models.Transaction{}, err
	}

	return &transaction, nil
}
