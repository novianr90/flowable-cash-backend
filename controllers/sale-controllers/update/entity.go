package updateSale

import "time"

type InputUpdateSale struct {
	ID          uint   `form:"id"`
	Name        string `form:"transaction_name"`
	Date        string `form:"transaction_date"`
	Total       uint   `form:"transaction_total"`
	FeeType     string `form:"fee_type"`
	Fee         uint   `form:"transaction_fee"`
	Description string `form:"description"`
}

type ResponseTransaction struct {
	ID          uint      `json:"id"`
	Date        string    `json:"date"`
	Name        string    `json:"name"`
	Total       uint      `json:"total"`
	Type        string    `json:"type"`
	FeeType     string    `json:"fee_type"`
	Fee         uint      `json:"transaction_fee"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
