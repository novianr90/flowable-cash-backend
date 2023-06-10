package createBalanceSheet

import (
	"flowable-cash-backend/helpers"
	"flowable-cash-backend/models"
)

type Service interface {
	CreateBalanceSheet(input *InputCreateBalanceSheet) (*ResponseBalanceSheet, error)
}

type service struct {
	repo Repository
}

func NewCreateBalanceSheetService(repo Repository) *service {
	return &service{repo: repo}
}

func (s *service) CreateBalanceSheet(input *InputCreateBalanceSheet) (*ResponseBalanceSheet, error) {

	query := models.BalanceSheet{
		AccountNo:   helpers.AccountNoDecider(input.AccountName),
		AccountName: input.AccountName,
		Balance:     helpers.DebitCreditDecider(input.AccountName, input.Balance),
	}

	res, err := s.repo.CreateBalanceSheet(&query)

	if err != nil {
		return &ResponseBalanceSheet{}, err
	}

	response := ResponseBalanceSheet{
		ID:          res.ID,
		AccountNo:   res.AccountNo,
		AccountName: res.AccountName,
		Balance:     res.Balance,
		CreatedAt:   res.CreatedAt,
		UpdatedAt:   res.UpdatedAt,
	}

	return &response, nil

}
