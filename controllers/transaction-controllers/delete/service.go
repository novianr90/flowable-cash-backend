package deleteTransaction

import "flowable-cash-backend/models"

type Service interface {
	DeleteTransactionService(input *InputDeleteTransaction) (*models.Transaction, error)
}

type service struct {
	repo Repository
}

func NewDeleteService(repository Repository) *service {
	return &service{repo: repository}
}

func (s *service) DeleteTransactionService(input *InputDeleteTransaction) (*models.Transaction, error) {
	transaction := models.Transaction{
		ID: input.ID,
	}

	result, err := s.repo.DeleteTransactionRepository(&transaction)

	if err != nil {
		return &models.Transaction{}, err
	}

	return result, nil
}
