package readTransaction

import (
	"flowable-cash-backend/helpers"
)

type Service interface {
	ReadAllTransactions() ([]ResponseTransaction, []ResponseTransaction, error)
}

type service struct {
	repo Repository
}

func NewReadService(repo Repository) *service {
	return &service{repo: repo}
}

func (s *service) ReadAllTransactions() ([]ResponseTransaction, []ResponseTransaction, error) {

	pengeluaranRaw, pemasukkanRaw, err := s.repo.ReadAllTransactions()

	var pengeluaran []ResponseTransaction
	var pemasukkan []ResponseTransaction

	for _, value := range *pengeluaranRaw {

		formattedDate := helpers.DateToString(value.Date)

		pengeluaran = append(pengeluaran, ResponseTransaction{
			ID:          value.ID,
			Date:        formattedDate,
			Name:        value.Name,
			Total:       value.Total,
			Payment:     value.Payment,
			Description: value.Description,
			CreatedAt:   value.CreatedAt,
			UpdatedAt:   value.UpdatedAt,
		})
	}

	for _, value := range *pemasukkanRaw {

		formattedDate := helpers.DateToString(value.Date)

		pemasukkan = append(pemasukkan, ResponseTransaction{
			ID:          value.ID,
			Date:        formattedDate,
			Name:        value.Name,
			Total:       value.Total,
			Payment:     value.Payment,
			Description: value.Description,
			CreatedAt:   value.CreatedAt,
			UpdatedAt:   value.UpdatedAt,
		})
	}

	if err != nil {
		return []ResponseTransaction{}, []ResponseTransaction{}, err
	}

	return pengeluaran, pemasukkan, nil
}
