package readPosting

import (
	"flowable-cash-backend/models"

	"gorm.io/gorm"
)

type Repository interface {
	ReadAllPosting() (*[]models.Posting, error)
	ReadPostingByTrxIDAndAccountID(input *models.Posting) (*models.Posting, error)
}

type repository struct {
	db *gorm.DB
}

func NewReadPostingRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) ReadAllPosting() (*[]models.Posting, error) {
	var result []models.Posting

	if err := r.db.Model(&models.Posting{}).Find(&result).Error; err != nil {
		return &[]models.Posting{}, err
	}

	return &result, nil
}

func (r *repository) ReadPostingByTrxIDAndAccountID(input *models.Posting) (*models.Posting, error) {
	var result models.Posting

	if err := r.db.Model(&models.Posting{}).
		Where("transaction_id = ?", input.TransactionID).
		Where("account_id = ?", input.AccountID).
		First(&result).Error; err != nil {
		return &models.Posting{}, err
	}

	return &result, nil
}
