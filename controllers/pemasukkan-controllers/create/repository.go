package createTransaction

import (
	"flowable-cash-backend/models"

	"gorm.io/gorm"
)

type Repository interface {
	CreateTransactionRepository(input *models.Pemasukkan) (*models.Pemasukkan, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryCreate(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) CreateTransactionRepository(input *models.Pemasukkan) (*models.Pemasukkan, error) {

	transaction := models.Pemasukkan{
		Name:          input.Name,
		Date:          input.Date,
		Total:         input.Total,
		Description:   input.Description,
		Payment:       input.Payment,
		AlreadyPosted: 0,
	}

	if err := r.db.Model(&models.Pemasukkan{}).Create(&transaction).Error; err != nil {
		return &models.Pemasukkan{}, err
	}

	return &transaction, nil
}
