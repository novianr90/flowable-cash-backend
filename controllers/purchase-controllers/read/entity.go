package readPurchase

import "time"

type InputReadPurchaseTransaction struct {
	ID uint `form:"id"`
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
