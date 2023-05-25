package services

import (
	"errors"
	"flowable-cash-backend/models"

	"gorm.io/gorm"
)

type TransactionService struct {
	DB *gorm.DB
}

func (r *TransactionService) Create(transaction models.Transaction) (models.Transaction, error) {
	if err := r.DB.Create(&transaction).Error; err != nil {
		return models.Transaction{}, err
	}
	return transaction, nil
}

func (r *TransactionService) Update(id uint, transaction models.Transaction) (models.Transaction, error) {

	result := r.DB.Model(models.Transaction{}).Where("id = ?", id).Updates(&transaction)

	if result.Error != nil {
		return models.Transaction{}, result.Error
	}

	if result.RowsAffected == 0 {
		return models.Transaction{}, errors.New("no data to be updated")
	}

	return transaction, nil
}

func (r *TransactionService) Delete(id uint) error {

	result := r.DB.Where("id = ?", id).Delete(models.Transaction{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("no data to be deleted")
	}

	return nil
}

func (r *TransactionService) GetTransactions() ([]models.Transaction, error) {
	var transactions []models.Transaction

	if err := r.DB.Find(&transactions).Error; err != nil {
		return []models.Transaction{}, err
	}

	return transactions, nil
}

func (r *TransactionService) GetTransaction(id uint) (models.Transaction, error) {
	var transaction models.Transaction

	if err := r.DB.Where("id = ?", id).First(&transaction).Error; err != nil {
		return models.Transaction{}, err
	}

	return transaction, nil
}
