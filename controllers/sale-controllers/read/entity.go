package readSale

import "time"

type InputReadSaleTransaction struct {
	ID uint `form:"id"`
}

type ResponseTransaction struct {
	ID          uint      `json:"id"`
	Date        time.Time `json:"date"`
	Name        string    `json:"name"`
	Total       uint      `json:"total"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
