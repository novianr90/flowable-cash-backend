package createPosting

import "flowable-cash-backend/models"

type Service interface {
	CreateNewRecord(input *ModelCreatePosting) (*Response, error)
}

type service struct {
	repository Repository
}

func NewCreatePostingService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) CreateNewRecord(input *ModelCreatePosting) (*Response, error) {

	query := models.Posting{
		TransactionID: input.TransactionID,
		AccountID:     input.AccountID,
	}

	res, err := s.repository.CreateNewRecord(&query)

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
