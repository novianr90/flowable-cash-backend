package sorting

import (
	"flowable-cash-backend/models"

	"gorm.io/gorm"
)

type Sorting interface {
	SortTransaction() ([]models.Transaction, error)
}

type sorting struct {
	db *gorm.DB
}

func NewSortingInternal(db *gorm.DB) *sorting {
	return &sorting{db: db}
}

func (s *sorting) SortTransaction() ([]models.Transaction, error) {

	var (
		transactions []models.Transaction

		sale models.Sale

		purchase models.Purchase
	)

	db := s.db.Model(&models.Transaction{})

	err := db.Find(&transactions).Error

	if err != nil {
		return nil, err
	}

	for _, value := range transactions {

		if value.Type == "Sale" {
			sale = models.Sale{
				Name:          value.Name,
				Date:          value.Date,
				Total:         value.Total,
				Description:   value.Description,
				TransactionID: value.ID,
			}

			err := s.db.Model(&models.Sale{}).Create(&sale).Error

			if err != nil {
				return nil, err
			}

		} else {

			purchase = models.Purchase{
				Name:          value.Name,
				Date:          value.Date,
				Total:         value.Total,
				Description:   value.Description,
				TransactionID: value.ID,
			}

			err := s.db.Model(&models.Purchase{}).Create(&purchase).Error

			if err != nil {
				return nil, err
			}

		}

	}

	return transactions, nil
}
