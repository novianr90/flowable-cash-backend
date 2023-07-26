package createBalanceSheet

import (
	"encoding/json"
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

	balance, err := json.Marshal(input.Balance)

	if err != nil {
		return &ResponseBalanceSheet{}, err
	}

	query := models.BalanceSheet{
		AccountNo:   helpers.AccountNoDecider(input.AccountName),
		AccountName: input.AccountName,
		Balance:     balance,
		Month:       input.Month,
	}

	res, err := s.repo.CreateBalanceSheet(&query)

	if err != nil {
		return &ResponseBalanceSheet{}, err
	}

	var balanceRes models.Balance

	_ = json.Unmarshal(res.Balance, &balanceRes)

	response := ResponseBalanceSheet{
		ID:          res.ID,
		AccountNo:   res.AccountNo,
		AccountName: res.AccountName,
		Month:       res.Month,
		Balance:     balanceRes,
		CreatedAt:   res.CreatedAt,
		UpdatedAt:   res.UpdatedAt,
	}

	return &response, nil

}
