package updateTransaction

import "time"

type InputUpdateTransaction struct {
	ID            uint   `form:"id"`
	Name          string `form:"transaction_name"`
	Date          string `form:"transaction_date"`
	Type          string `form:"transaction_type"`
	Total         uint   `form:"transaction_total"`
	Description   string `form:"description"`
	FeeType       string `form:"fee_type"`
	Fee           uint   `form:"transaction_fee"`
	AlreadyPosted uint   `form:"already_posted"`
}

type ResponseTransaction struct {
	ID            uint      `json:"id"`
	Date          string    `json:"date"`
	Name          string    `json:"name"`
	Type          string    `json:"type"`
	Total         uint      `json:"total"`
	FeeType       string    `json:"fee_type"`
	Fee           uint      `json:"transaction_fee"`
	Description   string    `json:"description"`
	AlreadyPosted uint      `json:"already_posted"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
