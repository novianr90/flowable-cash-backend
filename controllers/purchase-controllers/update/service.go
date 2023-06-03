package updatePurchase

import (
	"flowable-cash-backend/helpers"
	"flowable-cash-backend/models"
)

type Service interface {
	UpdatePurchaseTransactionById(input *InputUpdatePurchase) (ResponseTransaction, error)
}

type service struct {
	repo Repository
}

func NewServiceUpdatePurchase(repo Repository) *service {
	return &service{repo: repo}
}

func (s *service) UpdatePurchaseTransactionById(input *InputUpdatePurchase) (ResponseTransaction, error) {
	formattedDate, _ := helpers.StringToDate(input.Date)

	update := models.Transaction{
		ID:          input.ID,
		Name:        input.Name,
		Date:        formattedDate,
		Total:       input.Total,
		Description: input.Description,
	}

	res, err := s.repo.UpdatePurchaseTransaction(&update)

	response := ResponseTransaction{
		ID:          res.ID,
		Name:        res.Name,
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
