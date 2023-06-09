package readSale

import "time"

type InputReadSaleTransaction struct {
	ID uint `form:"id"`
}

type ResponseTransaction struct {
	ID          uint      `json:"id"`
	Date        string    `json:"date"`
	Name        string    `json:"name"`
	Type        string    `json:"type"`
	Total       uint      `json:"total"`
	FeeType     string    `json:"fee_type"`
	Fee         uint      `json:"transaction_fee"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
