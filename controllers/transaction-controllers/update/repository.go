package updateTransaction

import (
	"flowable-cash-backend/internal/sorting"
	"flowable-cash-backend/models"

	"gorm.io/gorm"
)

type Repository interface {
	UpdateTransactionRepository(input *models.Transaction) (*models.Transaction, error)
}

type repository struct {
	db      *gorm.DB
	sorting sorting.Sorting
}

func NewRepositoryUpdate(db *gorm.DB, sorting sorting.Sorting) *repository {
	return &repository{db: db, sorting: sorting}
}

func (r *repository) UpdateTransactionRepository(input *models.Transaction) (*models.Transaction, error) {

	var transaction models.Transaction

	db := r.db.Model(models.Transaction{})

	newTransaction := models.Transaction{
		Name:        input.Name,
		Date:        input.Date,
		Type:        input.Type,
		Total:       input.Total,
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

	if err := r.sorting.SortTransaction(); err != nil {
		return &models.Transaction{}, err
	}

	return &transaction, nil
}
