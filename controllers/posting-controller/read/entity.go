package readPosting

import "time"

type InputRead struct {
	TransactionID uint `json:"transaction_id" form:"transaction_id"`
	AccountID     uint `json:"account_id" form:"account_id"`
}

type Response struct {
	ID            uint      `json:"posting_id" form:"posting_id"`
	TransactionID uint      `json:"transaction_id" form:"transaction_id"`
	AccountID     uint      `json:"account_id" form:"account_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
