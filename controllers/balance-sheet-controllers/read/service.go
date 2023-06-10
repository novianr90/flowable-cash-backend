package readBalanceSheet

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

	res, err := s.repo.GetBalanceSheet()

	if err != nil {
		return &[]ResponseBalanceSheet{}, err
	}

	for _, value := range *res {
		response = append(response, ResponseBalanceSheet{
			ID:          value.ID,
			AccountNo:   value.AccountNo,
			AccountName: value.AccountName,
			Balance:     value.Balance,
			CreatedAt:   value.CreatedAt,
			UpdatedAt:   value.UpdatedAt,
		})
	}

	return &response, nil

}
