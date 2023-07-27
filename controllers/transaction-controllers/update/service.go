package updateTransaction

import (
	"errors"
	"flowable-cash-backend/helpers"
	"flowable-cash-backend/models"
)

type Service interface {
	UpdateTransactionService(input *InputUpdateTransaction) error
}

type service struct {
	repo Repository
}

func NewUpdateService(repo Repository) *service {
	return &service{repo: repo}
}

func (s *service) UpdateTransactionService(input *InputUpdateTransaction) error {

	trxType := input.Type

	if trxType == "Pemasukkan" {

		formattedDate, _ := helpers.StringToDate(input.Date)

		updatePemasukkan := models.Pemasukkan{
			ID:          input.ID,
			Description: input.Description,
			Total:       input.Total,
			Date:        formattedDate,
		}

		err := s.repo.UpdateTransactionRepository(input.Type, updatePemasukkan)

		if err != nil {
			return err
		}
	} else {

		formattedDate, _ := helpers.StringToDate(input.Date)

		updatePengeluaran := models.Pengeluaran{
			ID:          input.ID,
			Description: input.Description,
			Total:       input.Total,
			Date:        formattedDate,
		}

		err := s.repo.UpdateTransactionRepository(input.Type, updatePengeluaran)

		if err != nil {
			return err
		}

	}

	return errors.New("please input type")
}
