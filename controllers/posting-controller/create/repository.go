package createPosting

import (
	"flowable-cash-backend/models"

	"gorm.io/gorm"
)

type Repository interface {
	CreateNewRecord(input *models.Posting) (*models.Posting, error)
}

type repository struct {
	db *gorm.DB
}

func NewCreatePostingRepo(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) CreateNewRecord(input *models.Posting) (*models.Posting, error) {

	var posting models.Posting

	query := models.Posting{
		TransactionID: input.TransactionID,
		AccountID:     input.AccountID,
	}

	result := r.db.Model(&models.Posting{}).Save(&query)

	if result.Error != nil {
		return &models.Posting{}, result.Error
	}

	err := r.db.Model(&models.Posting{}).Where("transaction_id = ?", input.TransactionID).First(&posting).Error

	if err != nil {
		return &models.Posting{}, err
	}

	return &posting, nil
}
