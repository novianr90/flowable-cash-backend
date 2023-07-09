package readTransaction

import (
	"flowable-cash-backend/helpers"
	"flowable-cash-backend/models"
)

type Service interface {
	ReadAllTransactions() ([]ResponseTransaction, error)
	ReadTransactionById(input *InputReadTransaction) (ResponseTransaction, error)
	ReadAllTransactionsByType(input *InputReadTransaction) ([]ResponseTransaction, error)
}

type service struct {
	repo Repository
}

func NewReadService(repo Repository) *service {
	return &service{repo: repo}
}

func (s *service) ReadAllTransactions() ([]ResponseTransaction, error) {
	result, err := s.repo.ReadAllTransactions()

	var responses []ResponseTransaction

	for _, value := range *result {

		formattedDate := helpers.DateToString(value.Date)

		responses = append(responses, ResponseTransaction{
			ID:          value.ID,
			Date:        formattedDate,
			Name:        value.Name,
			Type:        value.Type,
			Total:       value.Total,
			FeeType:     value.FeeType,
			Fee:         value.Fee,
			Description: value.Description,
			CreatedAt:   value.CreatedAt,
			UpdatedAt:   value.UpdatedAt,
		})
	}

	if err != nil {
		return []ResponseTransaction{}, err
	}

	return responses, nil
}

func (s *service) ReadTransactionById(input *InputReadTransaction) (ResponseTransaction, error) {

	transaction := models.Transaction{
		ID: input.ID,
	}

	result, err := s.repo.ReadTransactionById(&transaction)

	formattedDate := helpers.DateToString(result.Date)

	response := ResponseTransaction{
		ID:          result.ID,
		Date:        formattedDate,
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

func (s *service) ReadAllTransactionsByType(input *InputReadTransaction) ([]ResponseTransaction, error) {
	transactions := models.Transaction{
		Type: input.Type,
	}

	res, err := s.repo.ReadAllTransactionsByType(&transactions)

	if err != nil {
		return []ResponseTransaction{}, err
	}

	var responses []ResponseTransaction

	for _, value := range *res {

		formattedDate := helpers.DateToString(value.Date)

		responses = append(responses, ResponseTransaction{
			ID:          value.ID,
			Date:        formattedDate,
			Name:        value.Name,
			Type:        value.Type,
			Total:       value.Total,
			FeeType:     value.FeeType,
			Fee:         value.Fee,
			Description: value.Description,
			CreatedAt:   value.CreatedAt,
			UpdatedAt:   value.UpdatedAt,
		})
	}

	return responses, nil
}
