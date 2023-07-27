package deleteTransaction

import (
	"errors"
	"flowable-cash-backend/models"
)

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

	var err error

	switch {

	case input.Type == "Pemasukkan":
		queryPemasukkan := models.Pemasukkan{ID: input.ID}
		err = s.repo.DeletePemasukkan(&queryPemasukkan)

	case input.Type == "Pengeluaran":
		queryPengeluaran := models.Pengeluaran{ID: input.ID}
		err = s.repo.DeletePengeluaran(&queryPengeluaran)

	default:
		return errors.New("specify transaction type")
	}

	if err != nil {
		return err
	}

	return nil
}
