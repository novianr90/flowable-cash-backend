package readPosting

import "flowable-cash-backend/models"

type Service interface {
	ReadAllPosting() (*[]Response, error)
	ReadPostingByTrxIDAndAccountID(input *InputRead) (*Response, error)
}

type service struct {
	repository Repository
}

func NewReadPostingService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) ReadAllPosting() (*[]Response, error) {
	var response []Response

	res, err := s.repository.ReadAllPosting()

	if err != nil {
		return &[]Response{}, err
	}

	for _, value := range *res {
		response = append(response, Response{
			ID:            value.ID,
			AccountID:     value.AccountID,
			TransactionID: value.TransactionID,
			CreatedAt:     value.CreatedAt,
			UpdatedAt:     value.UpdatedAt,
		})
	}

	return &response, nil
}

func (s *service) ReadPostingByTrxIDAndAccountID(input *InputRead) (*Response, error) {
	query := models.Posting{
		TransactionID: input.TransactionID,
		AccountID:     input.AccountID,
	}

	res, err := s.repository.ReadPostingByTrxIDAndAccountID(&query)

	if err != nil {
		return &Response{}, err
	}

	response := Response{
		ID:            res.ID,
		TransactionID: res.TransactionID,
		AccountID:     res.AccountID,
		CreatedAt:     res.CreatedAt,
		UpdatedAt:     res.UpdatedAt,
	}

	return &response, nil
}
