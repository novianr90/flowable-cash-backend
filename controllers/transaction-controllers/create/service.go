package createTransaction

import (
	"flowable-cash-backend/helpers"
	"flowable-cash-backend/models"
)

type Service interface {
	CreateTransactionService(input *InputCreateTransaction) (ResponseTransaction, error)
}

type service struct {
	repo Repository
}

func NewServiceCreate(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateTransactionService(input *InputCreateTransaction) (ResponseTransaction, error) {

	formattedDate, _ := helpers.StringToDate(input.Date)

	transaction := models.Transaction{
		Name:        input.Name,
		Date:        formattedDate,
		Type:        input.Type,
		Total:       input.Total,
		Description: input.Description,
		FeeType:     input.FeeType,
		Fee:         input.Fee,
	}

	result, err := s.repo.CreateTransactionRepository(&transaction)

	date := helpers.DateToString(result.Date)

	response := ResponseTransaction{
		ID:          result.ID,
		Date:        date,
		Name:        result.Name,
		Type:        result.Type,
		Total:       result.Total,
		FeeType:     result.FeeType,
		Fee:         result.Fee,
		Description: result.Description,
		CreatedAt:   result.CreatedAt,
		UpdatedAt:   result.UpdatedAt,
	}

	if err != nil {
		return ResponseTransaction{}, err
	}

	return response, nil
}
