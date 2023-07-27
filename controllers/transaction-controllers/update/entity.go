package updateTransaction

import "time"

type InputUpdateTransaction struct {
	ID          uint   `form:"id"`
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
	Total       uint      `json:"total"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
