package updatePurchase

import (
	"errors"
	"flowable-cash-backend/models"

	"gorm.io/gorm"
)

type Repository interface {
	UpdatePurchaseTransaction(input *models.Transaction) (*models.Transaction, error)
}

type repository struct {
	db *gorm.DB
}

func NewUpdatePurchaseRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) UpdatePurchaseTransaction(input *models.Transaction) (*models.Transaction, error) {
	typePurchase := "Purchase"

	db := r.db.Model(&models.Transaction{}).Where("type = ?", typePurchase)

	update := models.Transaction{
		ID:          input.ID,
		Name:        input.Name,
		Date:        input.Date,
		Total:       input.Total,
		Description: input.Description,
	}

	res := db.Where("id = ?", input.ID).Updates(&update)

	if res.Error != nil {
		return &models.Transaction{}, res.Error
	}

	if res.RowsAffected == 0 {
		return &models.Transaction{}, errors.New("no data found")
	}

	return &update, nil
}
