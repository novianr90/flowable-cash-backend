package readTransaction

import "time"

type ResponseTransaction struct {
	ID          uint      `json:"id"`
	Date        string    `json:"date"`
	Name        string    `json:"name"`
	Total       uint      `json:"total"`
	Payment     string    `json:"transaction_payment"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
