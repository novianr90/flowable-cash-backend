package readSale

import (
	"flowable-cash-backend/helpers"
	"flowable-cash-backend/models"
)

type Service interface {
	ReadAllSaleTypeTransactions() ([]ResponseTransaction, error)
	ReadSaleTypeById(input *InputReadSaleTransaction) (ResponseTransaction, error)
}

type service struct {
	repo Repository
}

func NewReadSaleService(repo Repository) *service {
	return &service{repo: repo}
}

func (s *service) ReadAllSaleTypeTransactions() ([]ResponseTransaction, error) {
	result, err := s.repo.ReadAllSaleTypeTransactions()

	var responses []ResponseTransaction

	for _, value := range *result {

		formattedDate := helpers.DateToString(value.Date)

		responses = append(responses, ResponseTransaction{
			ID:          value.ID,
			Date:        formattedDate,
			Name:        value.Name,
			Total:       value.Total,
			Description: value.Description,
			CreatedAt:   value.CreatedAt,
			UpdatedAt:   value.UpdatedAt,
			FeeType:     value.FeeType,
			Fee:         value.Fee,
		})
	}

	if err != nil {
		return []ResponseTransaction{}, err
	}

	return responses, nil
}

func (s *service) ReadSaleTypeById(input *InputReadSaleTransaction) (ResponseTransaction, error) {

	query := models.Transaction{
		ID: input.ID,
	}

	result, err := s.repo.ReadSaleTypeById(&query)

	formattedDate := helpers.DateToString(result.Date)

	response := ResponseTransaction{
		ID:          result.ID,
		Date:        formattedDate,
		Name:        result.Name,
		Total:       result.Total,
		Description: result.Description,
		CreatedAt:   result.CreatedAt,
		UpdatedAt:   result.UpdatedAt,
		Fee:         result.Fee,
		FeeType:     result.FeeType,
	}

	if err != nil {
		return ResponseTransaction{}, err
	}

	return response, nil
}
