package readPurchase

import "flowable-cash-backend/models"

type Service interface {
	// Read: Tx Purchase Type
	ReadAllPurchaseTypeTransactions() ([]ResponseTransaction, error)
	ReadPurchaseTypeById(input *InputReadPurchaseTransaction) (ResponseTransaction, error)
}

type service struct {
	repository Repository
}

func NewReadPurchaseService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) ReadAllPurchaseTypeTransactions() ([]ResponseTransaction, error) {
	result, err := s.repository.ReadAllPurchaseTypeTransactions()

	var responses []ResponseTransaction

	for _, value := range *result {
		responses = append(responses, ResponseTransaction{
			ID:          value.ID,
			Date:        value.Date,
			Name:        value.Name,
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

func (s *service) ReadPurchaseTypeById(input *InputReadPurchaseTransaction) (ResponseTransaction, error) {

	query := models.Transaction{
		ID: input.ID,
	}

	result, err := s.repository.ReadPurchaseTypeById(&query)

	response := ResponseTransaction{
		ID:          result.ID,
		Date:        result.Date,
		Name:        result.Name,
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
