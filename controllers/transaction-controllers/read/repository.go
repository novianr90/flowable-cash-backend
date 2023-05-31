package readTransaction

import (
	"flowable-cash-backend/models"

	"gorm.io/gorm"
)

type Repository interface {
	ReadAllTransactions() (*[]models.Transaction, error)
	ReadTransactionById(input *models.Transaction) (*models.Transaction, error)

	// Read: Tx Sale Type
	ReadAllSaleTypeTransactions() (*[]models.Transaction, error)
	ReadSaleTypeById(input *models.Transaction) (*models.Transaction, error)

	// Read: Tx Purchase Type
	ReadAllPurchaseTypeTransactions() (*[]models.Transaction, error)
	ReadPurchaseTypeById(input *models.Transaction) (*models.Transaction, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryRead(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) ReadAllTransactions() (*[]models.Transaction, error) {
	var transactions []models.Transaction

	db := r.db.Model(&transactions)

	result := db.Find(&transactions)

	if result.Error != nil {
		return &transactions, result.Error
	}

	return &transactions, nil
}

func (r *repository) ReadTransactionById(input *models.Transaction) (*models.Transaction, error) {
	var transaction models.Transaction

	db := r.db.Model(&transaction)

	result := db.Where("id = ?", input.ID).First(&transaction)

	if result.Error != nil {
		return &transaction, result.Error
	}

	return &transaction, nil
}

func (r *repository) ReadAllSaleTypeTransactions() (*[]models.Transaction, error) {
	typeSale := "Sale"

	var saleTransactions []models.Transaction

	db := r.db.Model(&models.Transaction{})

	err := db.Where("type = ?", typeSale).Find(&saleTransactions).Error

	if err != nil {
		return &[]models.Transaction{}, err
	}

	return &saleTransactions, nil
}

func (r *repository) ReadSaleTypeById(input *models.Transaction) (*models.Transaction, error) {
	typeSale := "Sale"

	var saleTransaction models.Transaction

	db := r.db.Model(&models.Transaction{})

	err := db.Where("type = ?", typeSale).Where("id = ?", input.ID).First(&saleTransaction).Error

	if err != nil {
		return &models.Transaction{}, err
	}

	return &saleTransaction, nil
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
