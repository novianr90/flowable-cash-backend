package routes

import (
	deleteTransaction "flowable-cash-backend/controllers/transaction-controllers/delete"
	readTransaction "flowable-cash-backend/controllers/transaction-controllers/read"
	updateTransaction "flowable-cash-backend/controllers/transaction-controllers/update"

	handlerDelete "flowable-cash-backend/handlers/transaction-handlers/delete"
	handlerRead "flowable-cash-backend/handlers/transaction-handlers/read"
	handlerUpdate "flowable-cash-backend/handlers/transaction-handlers/update"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

func InitTransactionRoutes(db *gorm.DB, route *gin.RouterGroup) {

	//Delete
	deleteRepository := deleteTransaction.NewRepositoryDelete(db)
	deleteService := deleteTransaction.NewDeleteService(deleteRepository)
	deleteHandler := handlerDelete.NewHandlerDeleteTransaction(deleteService)

	// Read
	readRepository := readTransaction.NewRepositoryRead(db)
	readService := readTransaction.NewReadService(readRepository)
	readHandler := handlerRead.NewHandlerReadTransaction(readService)

	// Update
	updateRepository := updateTransaction.NewRepositoryUpdate(db)
	updateService := updateTransaction.NewUpdateService(updateRepository)
	updateHandler := handlerUpdate.NewHandlerUpdateTransaction(updateService)

	groupRoute := route.Group("/transactions")

	groupRoute.DELETE("", deleteHandler.DeleteTransaction)

	// All transactions
	groupRoute.GET("", readHandler.GetAllTransactions)

	groupRoute.PUT("", updateHandler.UpdateTransaction)
}
