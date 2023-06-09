package deleteTransaction

import (
	"errors"
	"flowable-cash-backend/models"

	"gorm.io/gorm"
)

type Repository interface {
	DeleteTransactionRepository(input *models.Transaction) (*models.Transaction, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryDelete(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) DeleteTransactionRepository(input *models.Transaction) (*models.Transaction, error) {

	var transactionModel models.Transaction

	db := r.db.Model(&transactionModel)

	checkTransaction := db.Select("*").Where("id = ?", input.ID).Find(&transactionModel)

	if checkTransaction.RowsAffected < 1 {
		return &transactionModel, errors.New("no data found")
	}

	deleteStudentId := db.Select("*").Where("id = ?", input.ID).Find(&transactionModel).Delete(&transactionModel)

	if deleteStudentId.Error != nil {
		return &transactionModel, deleteStudentId.Error
	}

	return &transactionModel, nil
}
