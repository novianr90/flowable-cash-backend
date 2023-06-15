package updateBalanceSheet

import (
	"flowable-cash-backend/models"
	"time"
)

type InputUpdateBalanceSheet struct {
	ID          uint   `form:"balance_sheet_id"`
	AccountName string `form:"account_name"`
	Balance     string `form:"account_balance"`
}

type ResponseBalanceSheet struct {
	ID          uint           `json:"balance_sheet_id"`
	AccountNo   string         `json:"account_no"`
	AccountName string         `json:"account_name"`
	Balance     models.Balance `json:"account_balance"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}
