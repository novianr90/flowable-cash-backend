package updateBalanceSheet

import (
	"encoding/json"
	"flowable-cash-backend/helpers"
	"flowable-cash-backend/models"
	"strings"
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

	var balanceFormatted models.Balance

	formattedBalanceInString := input.Balance

	formattedBalanceInString = strings.ReplaceAll(formattedBalanceInString, "\\", "")
	formattedBalanceInString = strings.Replace(formattedBalanceInString, ",\n}", "\n}", 1)

	_ = json.Unmarshal([]byte(formattedBalanceInString), &balanceFormatted)

	balance, _ := json.Marshal(balanceFormatted)

	query := models.BalanceSheet{
		ID:          input.ID,
		AccountNo:   helpers.AccountNoDecider(input.AccountName),
		AccountName: input.AccountName,
		Balance:     balance,
	}

	res, err := s.repo.UpdateBalanceSheet(&query)

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
	}

	return response, nil

}
