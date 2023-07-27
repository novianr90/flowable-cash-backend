package readPurchase

import (
	"flowable-cash-backend/models"

	"gorm.io/gorm"
)

type Repository interface {
	ReadAllPurchaseTypeTransactions() (*[]models.Pengeluaran, error)
	ReadPurchaseTypeById(input *models.Pengeluaran) (*models.Pengeluaran, error)
}

type repository struct {
	db *gorm.DB
}

func NewReadPurchaseRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

var typePurchase = "Pengeluaran"

func (r *repository) ReadAllPurchaseTypeTransactions() (*[]models.Pengeluaran, error) {

	var purchaseTransactions []models.Pengeluaran

	db := r.db.Model(&models.Pengeluaran{})

	err := db.Where("type = ?", typePurchase).Find(&purchaseTransactions).Error

	if err != nil {
		return &[]models.Pengeluaran{}, err
	}

	return &purchaseTransactions, nil
}

func (r *repository) ReadPurchaseTypeById(input *models.Pengeluaran) (*models.Pengeluaran, error) {

	var purchaseTransaction models.Pengeluaran

	db := r.db.Model(&models.Pengeluaran{})

	err := db.Where("type = ?", typePurchase).Where("id = ?", input.ID).First(&purchaseTransaction).Error

	if err != nil {
		return &models.Pengeluaran{}, err
	}

	return &purchaseTransaction, nil
}
