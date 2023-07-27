package deleteTransaction

type Service interface {
	DeleteTransactionService(input *InputDeleteTransaction) error
}

type service struct {
	repo Repository
}

func NewDeleteService(repository Repository) *service {
	return &service{repo: repository}
}

func (s *service) DeleteTransactionService(input *InputDeleteTransaction) error {

	err := s.repo.DeleteTransactionRepository(input.ID, input.Type)

	if err != nil {
		return err
	}

	return nil
}
