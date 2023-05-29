package updateTransaction

import (
	"flowable-cash-backend/helpers"
	"flowable-cash-backend/models"
)

type Service interface {
	UpdateTransactionService(input *InputUpdateTransaction) (*models.Transaction, error)
}

type service struct {
	repo Repository
}

func NewUpdateService(repo Repository) *service {
	return &service{repo: repo}
}

func (s *service) UpdateTransactionService(input *InputUpdateTransaction) (*models.Transaction, error) {

	formattedDate, _ := helpers.StringToDate(input.Date)

	transaction := models.Transaction{
		ID:          input.ID,
		Name:        input.Name,
		Date:        formattedDate,
		Type:        input.Type,
		Total:       input.Total,
		Description: input.Description,
	}

	result, err := s.repo.UpdateTransactionRepository(&transaction)

	if err != nil {
		return result, err
	}

	return result, nil
}
