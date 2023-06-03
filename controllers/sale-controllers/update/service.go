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
	}

	res, err := s.repo.UpdateSaleTransaction(&saleTransaction)

	response := ResponseTransaction{
		ID:          res.ID,
		Name:        res.Name,
		Date:        res.Date,
		Total:       res.Total,
		Description: res.Description,
		CreatedAt:   res.CreatedAt,
		UpdatedAt:   res.UpdatedAt,
	}

	if err != nil {
		return ResponseTransaction{}, err
	}

	return response, nil
}
