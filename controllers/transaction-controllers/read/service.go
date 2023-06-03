package readTransaction

import "flowable-cash-backend/models"

type Service interface {
	ReadAllTransactions() ([]ResponseTransaction, error)
	ReadTransactionById(input *InputReadTransaction) (ResponseTransaction, error)
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
		responses = append(responses, ResponseTransaction{
			ID:          value.ID,
			Date:        value.Date,
			Name:        value.Name,
			Type:        value.Type,
			Total:       value.Total,
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
