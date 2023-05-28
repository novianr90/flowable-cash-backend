package fields

import (
	"flowable-cash-backend/resolver"

	"github.com/graphql-go/graphql"
)

var transactionType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Transaction",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"date": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"type": &graphql.Field{
				Type: graphql.String,
			},
			"total": &graphql.Field{
				Type: graphql.Int,
			},
			"description": &graphql.Field{
				Type: graphql.String,
			},
			"created_at": &graphql.Field{
				Type: graphql.DateTime,
			},
			"updated_at": &graphql.Field{
				Type: graphql.DateTime,
			},
		},
	},
)

type TransactionField struct {
	Resolver *resolver.TransactionResolver
}

func (f *TransactionField) GetMutationField() graphql.Fields {
	var transactionFields = graphql.Fields{
		"createTransaction": &graphql.Field{
			Type:        transactionType,
			Description: "Create a transaction",
			Args: graphql.FieldConfigArgument{
				"date": &graphql.ArgumentConfig{
					Type:        graphql.NewNonNull(graphql.String),
					Description: "Select a date with format dd/MM/yyyy",
				},
				"name": &graphql.ArgumentConfig{
					Type:        graphql.NewNonNull(graphql.String),
					Description: "Input your transaction name",
				},
				"type": &graphql.ArgumentConfig{
					Type:        graphql.NewNonNull(graphql.String),
					Description: "Input your transaction type",
				},
				"total": &graphql.ArgumentConfig{
					Type:        graphql.NewNonNull(graphql.Int),
					Description: "Input your total transaction price",
				},
				"description": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Input your description",
				},
			},
			Resolve: f.Resolver.CreateTransaction,
		},

		"updateTransaction": &graphql.Field{
			Type:        transactionType,
			Description: "Update transaction by id",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type:        graphql.NewNonNull(graphql.Int),
					Description: "Select ID to update",
				},
				"date": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Select a date with format dd/MM/yyyy",
				},
				"name": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Input your transaction name",
				},
				"type": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Input your transaction type",
				},
				"total": &graphql.ArgumentConfig{
					Type:        graphql.Int,
					Description: "Input your total transaction price",
				},
				"description": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Input your description",
				},
			},
			Resolve: f.Resolver.UpdateTransaction,
		},

		"deleteTransaction": &graphql.Field{
			Type:        transactionType,
			Description: "Delete transation by id",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type:        graphql.NewNonNull(graphql.Int),
					Description: "Select ID to delete",
				},
			},
			Resolve: f.Resolver.DeleteTransaction,
		},
	}

	return transactionFields
}

func (f *TransactionField) GetQueryField() graphql.Fields {
	var transactionType = graphql.Fields{
		"getTransactions": &graphql.Field{
			Type:        graphql.NewList(transactionType),
			Description: "Get all transactions",
			Resolve:     f.Resolver.GetTransactions,
		},

		"getTransactionById": &graphql.Field{
			Type:        transactionType,
			Description: "Get transaction by id",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type:        graphql.NewNonNull(graphql.Int),
					Description: "Select specific ID to get",
				},
			},
			Resolve: f.Resolver.GetTransactionById,
		},
	}

	return transactionType
}
