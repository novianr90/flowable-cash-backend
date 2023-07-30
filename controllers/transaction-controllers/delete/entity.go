package deleteTransaction

type InputDeleteTransaction struct {
	ID   uint   `form:"id"`
	Type string `form:"transaction_type"`
}
