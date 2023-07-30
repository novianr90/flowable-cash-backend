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

	var err error

	switch {
	case input.Type == "Pemasukkan":
		formattedDate, _ := helpers.StringToDate(input.Date)
		updatePemasukkan := models.Pemasukkan{
			ID:          input.ID,
			Description: input.Description,
			Total:       input.Total,
			Date:        formattedDate,
		}
		err = s.repo.UpdatePemasukkan(&updatePemasukkan)

	case input.Type == "Pengeluaran":
		formattedDate, _ := helpers.StringToDate(input.Date)
		updatePengeluaran := models.Pengeluaran{
			ID:          input.ID,
			Description: input.Description,
			Total:       input.Total,
			Date:        formattedDate,
		}
		err = s.repo.UpdatePengeluaran(&updatePengeluaran)

	default:
		return errors.New("specify transaction type")
	}

	if err != nil {
		return err
	}

	return nil
}
