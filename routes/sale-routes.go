package routes

import (
	createSale "flowable-cash-backend/controllers/pemasukkan-controllers/create"
	readSale "flowable-cash-backend/controllers/pemasukkan-controllers/read"

	handlerCreateSale "flowable-cash-backend/handlers/pemasukkan-handlers/create"
	handlerReadSaleTransaction "flowable-cash-backend/handlers/pemasukkan-handlers/read"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitSaleTransactionRoutes(db *gorm.DB, route *gin.RouterGroup) {
	// Read
	readRepository := readSale.NewRepositoryRead(db)
	readService := readSale.NewReadSaleService(readRepository)
	readHandler := handlerReadSaleTransaction.NewReadSaleHandler(readService)

	// Create
	createRepo := createSale.NewRepositoryCreate(db)
	createService := createSale.NewServiceCreate(createRepo)
	createHandler := handlerCreateSale.NewHandlerCreateSale(createService)

	groupRoute := route.Group("/pemasukkan")

	// Read
	groupRoute.GET("", readHandler.GetAllSaleTransactions)
	groupRoute.GET("/", readHandler.GetSaleTransactionById)

	// Create
	groupRoute.POST("", createHandler.CreatePemasukkan)
}
