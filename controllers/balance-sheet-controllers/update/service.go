package updateBalanceSheet

import (
	"encoding/json"
	"flowable-cash-backend/models"
)

type Service interface {
	UpdateAccount(input *InputUpdateBalanceSheet) (ResponseBalanceSheet, error)
	UpdateAccountAdmin(input *InputUpdateBalanceSheet) error
}

type service struct {
	repo Repository
}

func NewUpdateBalanceSheetService(repo Repository) *service {
	return &service{repo: repo}
}

func (s *service) UpdateAccount(input *InputUpdateBalanceSheet) (ResponseBalanceSheet, error) {

	balance, _ := json.Marshal(input.Balance)

	query := models.BalanceSheet{
		ID:          input.ID,
		AccountName: input.AccountName,
		Balance:     balance,
		Month:       input.Month,
	}

	res, err := s.repo.UpdateAccount(&query)

	if err != nil {
		return ResponseBalanceSheet{}, err
	}

	var balanceRes models.Balance

	_ = json.Unmarshal(res.Balance, &balanceRes)

	response := ResponseBalanceSheet{
		ID:          res.ID,
		AccountName: res.AccountName,
		AccountNo:   res.AccountNo,
		Balance:     balanceRes,
		Month:       res.Month,
		CreatedAt:   res.CreatedAt,
		UpdatedAt:   res.UpdatedAt,
	}

	return response, nil

}

func (s *service) UpdateAccountAdmin(input *InputUpdateBalanceSheet) error {
	balance, _ := json.Marshal(input.Balance)

	query := models.BalanceSheet{
		AccountName: input.AccountName,
		Month:       input.Month,
		Balance:     balance,
	}

	err := s.repo.UpdateAccountAdmin(&query)

	if err != nil {
		return err
	}

	return nil
}
