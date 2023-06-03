package routes

import (
	readPurchase "flowable-cash-backend/controllers/purchase-controllers/read"
	updatePurchase "flowable-cash-backend/controllers/purchase-controllers/update"

	handlerReadPurchase "flowable-cash-backend/handlers/purchase-handlers/read"
	handlerUpdatePurchase "flowable-cash-backend/handlers/purchase-handlers/update"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitPurchaseRoutes(db *gorm.DB, routes *gin.RouterGroup) {
	// Read
	readRepo := readPurchase.NewReadPurchaseRepository(db)
	readService := readPurchase.NewReadPurchaseService(readRepo)
	readHandler := handlerReadPurchase.NewReadPurchaseHandler(readService)

	// Update
	updateRepo := updatePurchase.NewUpdatePurchaseRepository(db)
	updateService := updatePurchase.NewServiceUpdatePurchase(updateRepo)
	updateHandler := handlerUpdatePurchase.NewUpdatePurchaseHandler(updateService)

	groupRoute := routes.Group("/purchases")

	groupRoute.GET("", readHandler.GetPurchaseTransactions)
	groupRoute.GET("/", readHandler.GetPurchaseTransactionById)

	groupRoute.PUT("", updateHandler.UpdatePurchaseTransactionById)
}
