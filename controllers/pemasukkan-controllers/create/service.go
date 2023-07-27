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

	pemasukkan := models.Pemasukkan{
		Name:        input.Name,
		Date:        formattedDate,
		Total:       input.Total,
		Description: input.Description,
		Payment:     input.Payment,
	}

	result, err := s.repo.CreateTransactionRepository(&pemasukkan)

	date := helpers.DateToString(result.Date)

	response := ResponseTransaction{
		ID:          result.ID,
		Date:        date,
		Name:        result.Name,
		Total:       result.Total,
		Payment:     result.Payment,
		Description: result.Description,
		CreatedAt:   result.CreatedAt,
		UpdatedAt:   result.UpdatedAt,
	}

	if err != nil {
		return ResponseTransaction{}, err
	}

	return response, nil
}
