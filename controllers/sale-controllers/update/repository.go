package updateSale

import (
	"errors"
	"flowable-cash-backend/models"

	"gorm.io/gorm"
)

type Repository interface {
	UpdateSaleTransaction(input *models.Transaction) (*models.Transaction, error)
}

type repository struct {
	db *gorm.DB
}

func NewUpdateSaleTransactionRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) UpdateSaleTransaction(input *models.Transaction) (*models.Transaction, error) {

	typeSale := "Sale"

	db := r.db.Model(&models.Transaction{}).Where("type = ?", typeSale)

	transaction := models.Transaction{
		Name:        input.Name,
		Date:        input.Date,
		Total:       input.Total,
		Description: input.Description,
	}

	res := db.Where("id = ?", input.ID).Updates(&transaction)

	if res.Error != nil {
		return &models.Transaction{}, res.Error
	}

	if res.RowsAffected == 0 {
		return &models.Transaction{}, errors.New("no record found")
	}

	return &transaction, nil
}
