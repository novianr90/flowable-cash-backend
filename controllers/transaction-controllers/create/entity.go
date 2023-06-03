package createTransaction

import "time"

type InputCreateTransaction struct {
	Name        string `form:"transaction_name"`
	Date        string `form:"transaction_date"`
	Type        string `form:"transaction_type"`
	Total       uint   `form:"transaction_total"`
	Description string `form:"description"`
}

type ResponseTransaction struct {
	ID          uint      `json:"id"`
	Date        string    `json:"date"`
	Name        string    `json:"name"`
	Type        string    `json:"type"`
	Total       uint      `json:"total"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
