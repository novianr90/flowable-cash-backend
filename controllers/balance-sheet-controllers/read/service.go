package readBalanceSheet

import (
	"encoding/json"
	"flowable-cash-backend/models"
)

type Service interface {
	GetAllAccounts(input *InputReadBalanceSheet) (*[]ResponseBalanceSheet, error)
	GetAllAccountsByAccountName(input *InputReadBalanceSheet) (*ResponseBalanceSheet, error)
	GetAllSpecificAccounts(input *InputReadBalanceSheet) (*[]ResponseBalanceSheet, error)
}

type service struct {
	repo Repository
}

func NewReadBalanceSheetService(repo Repository) *service {
	return &service{repo: repo}
}

func (s *service) GetAllAccounts(input *InputReadBalanceSheet) (*[]ResponseBalanceSheet, error) {

	var response []ResponseBalanceSheet

	query := models.BalanceSheet{
		Month: input.Month,
	}

	res, err := s.repo.GetAllAccounts(&query)

	if err != nil {
		return &[]ResponseBalanceSheet{}, err
	}

	for _, value := range *res {

		var balance models.Balance

		_ = json.Unmarshal(value.Balance, &balance)

		response = append(response, ResponseBalanceSheet{
			ID:          value.ID,
			AccountNo:   value.AccountNo,
			AccountName: value.AccountName,
			Balance:     balance,
			Month:       value.Month,
			CreatedAt:   value.CreatedAt,
			UpdatedAt:   value.UpdatedAt,
		})
	}

	return &response, nil

}

func (s *service) GetAllAccountsByAccountName(input *InputReadBalanceSheet) (*ResponseBalanceSheet, error) {

	query := models.BalanceSheet{
		AccountName: input.AccountName,
		Month:       input.Month,
	}

	res, err := s.repo.GetAllAccountsByAccountName(&query)

	if err != nil {
		return &ResponseBalanceSheet{}, err
	}

	var balance models.Balance
	_ = json.Unmarshal(res.Balance, &balance)

	response := ResponseBalanceSheet{
		ID:          res.ID,
		AccountNo:   res.AccountNo,
		AccountName: res.AccountName,
		Balance:     balance,
		Month:       res.Month,
		CreatedAt:   res.CreatedAt,
		UpdatedAt:   res.UpdatedAt,
	}

	return &response, nil
}

func (s *service) GetAllSpecificAccounts(input *InputReadBalanceSheet) (*[]ResponseBalanceSheet, error) {
	query := models.BalanceSheet{
		AccountName: input.AccountName,
	}

	var response []ResponseBalanceSheet

	res, err := s.repo.GetAllSpecificAccounts(&query)

	if err != nil {
		return &[]ResponseBalanceSheet{}, err
	}

	for _, value := range *res {

		var balance models.Balance

		_ = json.Unmarshal(value.Balance, &balance)

		response = append(response, ResponseBalanceSheet{
			ID:          value.ID,
			AccountNo:   value.AccountNo,
			AccountName: value.AccountName,
			Balance:     balance,
			Month:       value.Month,
			CreatedAt:   value.CreatedAt,
			UpdatedAt:   value.UpdatedAt,
		})
	}

	return &response, nil
}
