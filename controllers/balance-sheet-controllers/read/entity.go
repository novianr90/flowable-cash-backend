package readBalanceSheet

import (
	"flowable-cash-backend/models"
	"time"
)

type ResponseBalanceSheet struct {
	ID          uint           `json:"balance_sheet_id"`
	AccountNo   string         `json:"account_no"`
	AccountName string         `json:"account_name"`
	Balance     models.Balance `json:"account_balance"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}
