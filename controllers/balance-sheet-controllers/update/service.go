package updateBalanceSheet

import (
	"flowable-cash-backend/helpers"
	"flowable-cash-backend/models"
)

type Service interface {
	UpdateBalanceSheet(input *InputUpdateBalanceSheet) (ResponseBalanceSheet, error)
}

type service struct {
	repo Repository
}

func NewUpdateBalanceSheetService(repo Repository) *service {
	return &service{repo: repo}
}

func (s *service) UpdateBalanceSheet(input *InputUpdateBalanceSheet) (ResponseBalanceSheet, error) {

	query := models.BalanceSheet{
		ID:          input.ID,
		AccountNo:   helpers.AccountNoDecider(input.AccountName),
		AccountName: input.AccountName,
		Balance:     helpers.DebitCreditDecider(input.AccountName, input.Balance),
	}

	res, err := s.repo.UpdateBalanceSheet(&query)

	if err != nil {
		return ResponseBalanceSheet{}, err
	}

	response := ResponseBalanceSheet{
		ID:          res.ID,
		AccountName: res.AccountName,
		AccountNo:   res.AccountNo,
		Balance:     res.Balance,
	}

	return response, nil

}
