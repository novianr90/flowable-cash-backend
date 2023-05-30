package routes

import (
	createTransaction "flowable-cash-backend/controllers/transaction-controllers/create"
	deleteTransaction "flowable-cash-backend/controllers/transaction-controllers/delete"
	readTransaction "flowable-cash-backend/controllers/transaction-controllers/read"
	updateTransaction "flowable-cash-backend/controllers/transaction-controllers/update"
	"flowable-cash-backend/internal/sorting"

	handlerCreate "flowable-cash-backend/handlers/transaction-handlers/create"
	handlerDelete "flowable-cash-backend/handlers/transaction-handlers/delete"
	handlerRead "flowable-cash-backend/handlers/transaction-handlers/read"
	handlerUpdate "flowable-cash-backend/handlers/transaction-handlers/update"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

func InitTransactionRoutes(db *gorm.DB, route *gin.Engine) {

	sorting := sorting.NewSortingInternal(db)

	// Create
	createRepository := createTransaction.NewRepositoryCreate(db, sorting)
	createService := createTransaction.NewServiceCreate(createRepository)
	createHandler := handlerCreate.NewHandlerCreateTransaction(createService)

	//Delete
	deleteRepository := deleteTransaction.NewRepositoryDelete(db, sorting)
	deleteService := deleteTransaction.NewDeleteService(deleteRepository)
	deleteHandler := handlerDelete.NewHandlerDeleteTransaction(deleteService)

	// Read
	readRepository := readTransaction.NewRepositoryRead(db)
	readService := readTransaction.NewReadService(readRepository)
	readHandler := handlerRead.NewHandlerReadTransaction(readService)

	// Update
	updateRepository := updateTransaction.NewRepositoryUpdate(db, sorting)
	updateService := updateTransaction.NewUpdateService(updateRepository)
	updateHandler := handlerUpdate.NewHandlerUpdateTransaction(updateService)

	groupRoute := route.Group("/api/transactions")

	groupRoute.POST("", createHandler.CreateTransaction)

	groupRoute.DELETE("", deleteHandler.DeleteTransaction)

	groupRoute.GET("/all", readHandler.GetAllTransactions)
	groupRoute.GET("/id", readHandler.GetTransactionById)

	groupRoute.PUT("", updateHandler.UpdateTransaction)
}
