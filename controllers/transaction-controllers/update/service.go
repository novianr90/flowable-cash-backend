package updateTransaction

import (
	"flowable-cash-backend/helpers"
	"flowable-cash-backend/models"
)

type Service interface {
	UpdateTransactionService(input *InputUpdateTransaction) (ResponseTransaction, error)
}

type service struct {
	repo Repository
}

func NewUpdateService(repo Repository) *service {
	return &service{repo: repo}
}

func (s *service) UpdateTransactionService(input *InputUpdateTransaction) (ResponseTransaction, error) {

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

	response := ResponseTransaction{
		ID:          result.ID,
		Date:        result.Date,
		Name:        result.Name,
		Type:        result.Type,
		Total:       result.Total,
		Description: result.Description,
		CreatedAt:   result.CreatedAt,
		UpdatedAt:   result.UpdatedAt,
	}

	if err != nil {
		return ResponseTransaction{}, err
	}

	return response, nil
}
