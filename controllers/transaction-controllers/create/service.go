package createTransaction

import (
	"flowable-cash-backend/helpers"
	"flowable-cash-backend/models"
)

type Service interface {
	CreateTransactionService(input *InputCreateTransaction) (*models.Transaction, error)
}

type service struct {
	repo Repository
}

func NewServiceCreate(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateTransactionService(input *InputCreateTransaction) (*models.Transaction, error) {

	formattedDate, _ := helpers.StringToDate(input.Date)

	transaction := models.Transaction{
		Name:        input.Name,
		Date:        formattedDate,
		Type:        input.Type,
		Total:       input.Total,
		Description: input.Description,
	}

	resultCreateTransaction, err := s.repo.CreateTransactionRepository(&transaction)

	if err != nil {
		return &models.Transaction{}, err
	}

	return resultCreateTransaction, nil
}
