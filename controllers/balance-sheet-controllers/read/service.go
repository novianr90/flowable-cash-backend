package readBalanceSheet

import (
	"encoding/json"
	"flowable-cash-backend/models"
)

type Service interface {
	GetBalanceSheet() (*[]ResponseBalanceSheet, error)
}

type service struct {
	repo Repository
}

func NewReadBalanceSheetService(repo Repository) *service {
	return &service{repo: repo}
}

func (s *service) GetBalanceSheet() (*[]ResponseBalanceSheet, error) {

	var response []ResponseBalanceSheet

	var balance models.Balance

	res, err := s.repo.GetBalanceSheet()

	if err != nil {
		return &[]ResponseBalanceSheet{}, err
	}

	for _, value := range *res {

		_ = json.Unmarshal(value.Balance, &balance)

		response = append(response, ResponseBalanceSheet{
			ID:          value.ID,
			AccountNo:   value.AccountNo,
			AccountName: value.AccountName,
			Balance:     balance,
			CreatedAt:   value.CreatedAt,
			UpdatedAt:   value.UpdatedAt,
		})
	}

	return &response, nil

}
