package readPurchase

import (
	"flowable-cash-backend/models"

	"gorm.io/gorm"
)

type Repository interface {
	ReadAllPurchaseTypeTransactions() (*[]models.Transaction, error)
	ReadPurchaseTypeById(input *models.Transaction) (*models.Transaction, error)
}

type repository struct {
	db *gorm.DB
}

func NewReadPurchaseRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) ReadAllPurchaseTypeTransactions() (*[]models.Transaction, error) {
	typePurchase := "Purchase"

	var purchaseTransactions []models.Transaction

	db := r.db.Model(&models.Transaction{})

	err := db.Where("type = ?", typePurchase).Find(&purchaseTransactions).Error

	if err != nil {
		return &[]models.Transaction{}, err
	}

	return &purchaseTransactions, nil
}

func (r *repository) ReadPurchaseTypeById(input *models.Transaction) (*models.Transaction, error) {
	typePurchase := "Purchase"

	var purchaseTransaction models.Transaction

	db := r.db.Model(&models.Transaction{})

	err := db.Where("type = ?", typePurchase).Where("id = ?", input.ID).First(&purchaseTransaction).Error

	if err != nil {
		return &models.Transaction{}, err
	}

	return &purchaseTransaction, nil
}
