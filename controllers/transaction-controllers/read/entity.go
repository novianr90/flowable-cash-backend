package readTransaction

import "time"

type InputReadTransaction struct {
	ID   uint   `form:"id"`
	Type string `form:"type"`
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
