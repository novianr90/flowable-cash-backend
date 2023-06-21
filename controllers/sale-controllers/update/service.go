package updateSale

import (
	"flowable-cash-backend/helpers"
	"flowable-cash-backend/models"
)

type Service interface {
	UpdateSaleTransaction(input *InputUpdateSale) (ResponseTransaction, error)
}

type service struct {
	repo Repository
}

func NewUpdateSaleTransactionService(repo Repository) *service {
	return &service{repo: repo}
}

func (s *service) UpdateSaleTransaction(input *InputUpdateSale) (ResponseTransaction, error) {
	formattedDate, _ := helpers.StringToDate(input.Date)

	saleTransaction := models.Transaction{
		ID:          input.ID,
		Name:        input.Name,
		Date:        formattedDate,
		Total:       input.Total,
		Description: input.Description,
		FeeType:     input.FeeType,
		Fee:         input.Fee,
	}

	res, err := s.repo.UpdateSaleTransaction(&saleTransaction)

	date := helpers.DateToString(res.Date)

	response := ResponseTransaction{
		ID:          res.ID,
		Name:        res.Name,
		Date:        date,
		Total:       res.Total,
		Description: res.Description,
		CreatedAt:   res.CreatedAt,
		UpdatedAt:   res.UpdatedAt,
		FeeType:     res.FeeType,
		Fee:         res.Fee,
	}

	if err != nil {
		return ResponseTransaction{}, err
	}

	return response, nil
}
