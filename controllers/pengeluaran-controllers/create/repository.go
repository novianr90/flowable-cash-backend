package createTransaction

import (
	"flowable-cash-backend/models"

	"gorm.io/gorm"
)

type Repository interface {
	CreateTransactionRepository(input *models.Pengeluaran) (*models.Pengeluaran, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryCreate(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) CreateTransactionRepository(input *models.Pengeluaran) (*models.Pengeluaran, error) {

	query := models.Pengeluaran{
		Name:          input.Name,
		Date:          input.Date,
		Total:         input.Total,
		Description:   input.Description,
		Payment:       input.Payment,
		AlreadyPosted: 0,
	}

	if err := r.db.Model(&models.Pengeluaran{}).Create(&query).Error; err != nil {
		return &models.Pengeluaran{}, err
	}

	return &query, nil
}
