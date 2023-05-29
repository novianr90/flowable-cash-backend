package readTransaction

import "flowable-cash-backend/models"

type Service interface {
	ReadAllTransactions() (*[]models.Transaction, error)
	ReadTransactionById(input *InputReadTransaction) (*models.Transaction, error)
}

type service struct {
	repo Repository
}

func NewReadService(repo Repository) *service {
	return &service{repo: repo}
}

func (s *service) ReadAllTransactions() (*[]models.Transaction, error) {
	result, err := s.repo.ReadAllTransactions()

	if err != nil {
		return &[]models.Transaction{}, err
	}

	return result, nil
}

func (s *service) ReadTransactionById(input *InputReadTransaction) (*models.Transaction, error) {

	transaction := models.Transaction{
		ID: input.ID,
	}

	result, err := s.repo.ReadTransactionById(&transaction)

	if err != nil {
		return &models.Transaction{}, err
	}

	return result, nil
}
