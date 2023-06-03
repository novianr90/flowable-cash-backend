package routes

import (
	readSale "flowable-cash-backend/controllers/sale-controllers/read"
	updateSale "flowable-cash-backend/controllers/sale-controllers/update"

	handlerReadSaleTransaction "flowable-cash-backend/handlers/sale-handlers/read"
	handlerUpdateTransaction "flowable-cash-backend/handlers/sale-handlers/update"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitSaleTransactionRoutes(db *gorm.DB, route *gin.RouterGroup) {
	// Read
	readRepository := readSale.NewRepositoryRead(db)
	readService := readSale.NewReadSaleService(readRepository)
	readHandler := handlerReadSaleTransaction.NewReadSaleHandler(readService)

	// Update
	updateRepository := updateSale.NewUpdateSaleTransactionRepository(db)
	updateService := updateSale.NewUpdateSaleTransactionService(updateRepository)
	updateHandler := handlerUpdateTransaction.NewUpdateSaleTransactionHandler(updateService)

	groupRoute := route.Group("/sales")

	// Read
	groupRoute.GET("", readHandler.GetAllSaleTransactions)
	groupRoute.GET("/", readHandler.GetSaleTransactionById)

	// Update
	groupRoute.PUT("", updateHandler.UpdateSaleTransactionById)
}
