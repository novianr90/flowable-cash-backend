package resolver

import (
	"flowable-cash-backend/helpers"
	"flowable-cash-backend/models"
	"flowable-cash-backend/services"

	"github.com/graphql-go/graphql"
)

type TransactionResolver struct {
	Service *services.TransactionService
}

func (r *TransactionResolver) CreateTransaction(p graphql.ResolveParams) (interface{}, error) {
	transactionDate, _ := p.Args["date"].(string)
	transactionName, _ := p.Args["name"].(string)
	transactionType, _ := p.Args["type"].(string)
	transactionTotal, _ := p.Args["total"].(int)
	transactionDescription, _ := p.Args["description"].(string)

	formatteDate, _ := helpers.DateFormatter(transactionDate)

	transaction := models.Transaction{
		Date:        formatteDate,
		Name:        transactionName,
		Type:        transactionType,
		Total:       uint(transactionTotal),
		Description: transactionDescription,
	}

	result, err := r.Service.Create(transaction)

	if err != nil {
		return "data not found", err
	}

	return result, nil
}

func (r *TransactionResolver) UpdateTransaction(p graphql.ResolveParams) (interface{}, error) {
	id, _ := p.Args["id"].(int)
	transactionDate, _ := p.Args["date"].(string)
	transactionName, _ := p.Args["name"].(string)
	transactionType, _ := p.Args["type"].(string)
	transactionTotal, _ := p.Args["total"].(int)
	transactionDescription, _ := p.Args["description"].(string)

	formattedDate, _ := helpers.DateFormatter(transactionDate)

	transaction := models.Transaction{
		Date:        formattedDate,
		Name:        transactionName,
		Type:        transactionType,
		Total:       uint(transactionTotal),
		Description: transactionDescription,
	}

	result, err := r.Service.Update(uint(id), transaction)

	if err != nil {
		return "data not found", err
	}

	return result, nil
}

func (r *TransactionResolver) DeleteTransaction(p graphql.ResolveParams) (interface{}, error) {
	id, _ := p.Args["id"].(int)

	err := r.Service.Delete(uint(id))

	if err != nil {
		return "data not found", err
	}

	return "Data sucessfully deleted", nil
}

func (r *TransactionResolver) GetTransactions(p graphql.ResolveParams) (interface{}, error) {
	transactions, err := r.Service.GetTransactions()

	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (r *TransactionResolver) GetTransactionById(p graphql.ResolveParams) (interface{}, error) {
	id, _ := p.Args["id"].(int)

	transaction, err := r.Service.GetTransaction(uint(id))

	if err != nil {
		return nil, err
	}

	return transaction, nil
}
