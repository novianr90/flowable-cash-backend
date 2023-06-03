package readSale

import "flowable-cash-backend/models"

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

func (s *service) ReadSaleTypeById(input *InputReadSaleTransaction) (ResponseTransaction, error) {

	query := models.Transaction{
		ID: input.ID,
	}

	result, err := s.repo.ReadSaleTypeById(&query)

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
