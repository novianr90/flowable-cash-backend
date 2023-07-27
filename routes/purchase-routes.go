package routes

import (
	createPurchase "flowable-cash-backend/controllers/pengeluaran-controllers/create"
	readPurchase "flowable-cash-backend/controllers/pengeluaran-controllers/read"

	handlerCreatePurchase "flowable-cash-backend/handlers/pengeluaran-handlers/create"
	handlerReadPurchase "flowable-cash-backend/handlers/pengeluaran-handlers/read"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitPurchaseRoutes(db *gorm.DB, routes *gin.RouterGroup) {
	// Read
	readRepo := readPurchase.NewReadPurchaseRepository(db)
	readService := readPurchase.NewReadPurchaseService(readRepo)
	readHandler := handlerReadPurchase.NewReadPurchaseHandler(readService)

	// Create
	createRepo := createPurchase.NewRepositoryCreate(db)
	createService := createPurchase.NewServiceCreate(createRepo)
	createHandler := handlerCreatePurchase.NewHandlerCreatePurchase(createService)

	groupRoute := routes.Group("/pengeluaran")

	groupRoute.GET("", readHandler.GetPurchaseTransactions)
	groupRoute.GET("/", readHandler.GetPurchaseTransactionById)

	groupRoute.POST("", createHandler.CreatePemasukkan)
}
