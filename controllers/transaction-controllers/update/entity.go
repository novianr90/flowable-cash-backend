package updateTransaction

type InputUpdateTransaction struct {
	ID          uint   `form:"id"`
	Name        string `form:"transaction_name"`
	Date        string `form:"transaction_date"`
	Type        string `form:"transaction_type"`
	Total       uint   `form:"transaction_total"`
	Description string `form:"description"`
}
