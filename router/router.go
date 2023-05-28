package router

import (
	field "flowable-cash-backend/fields"
	"flowable-cash-backend/resolver"
	"flowable-cash-backend/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	"gorm.io/gorm"
)

type Query struct {
	Query string `json:"query"`
}

func StartServer(db *gorm.DB) *gin.Engine {
	app := gin.Default()

	transactionService := services.TransactionService{
		DB: db,
	}

	transactionResolver := resolver.TransactionResolver{
		Service: &transactionService,
	}

	transactionField := field.TransactionField{
		Resolver: &transactionResolver,
	}

	app.Use(gin.Recovery())
	app.Use(gin.Logger())

	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name:   "Query",
		Fields: transactionField.GetQueryField(),
	})

	mutationType := graphql.NewObject(graphql.ObjectConfig{
		Name:   "Mutation",
		Fields: transactionField.GetMutationField(),
	})

	schemaConfig := graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	}

	schema, err := graphql.NewSchema(schemaConfig)

	if err != nil {
		log.Fatalf("Failed to create schema, error: %v", err)
	}

	app.POST("/graphql", func(c *gin.Context) {
		var query Query

		if err := c.ShouldBindJSON(&query); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		result := graphql.Do(graphql.Params{
			Schema:        schema,
			RequestString: query.Query,
		})

		if len(result.Errors) > 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, result)
	})

	return app
}
