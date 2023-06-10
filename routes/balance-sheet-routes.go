package routes

import (
	createBalanceSheet "flowable-cash-backend/controllers/balance-sheet-controllers/create"
	readBalanceSheet "flowable-cash-backend/controllers/balance-sheet-controllers/read"
	updateBalanceSheet "flowable-cash-backend/controllers/balance-sheet-controllers/update"

	handlerCreateBalanceSheet "flowable-cash-backend/handlers/balance-sheet-handlers/create"
	handlerReadBalanceSheet "flowable-cash-backend/handlers/balance-sheet-handlers/read"
	handlerUpdateBalanceSheet "flowable-cash-backend/handlers/balance-sheet-handlers/update"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitBalanceSheetRoutes(db *gorm.DB, routes *gin.RouterGroup) {
	// Create
	createRepo := createBalanceSheet.NewCreateBalanceSheetRepository(db)
	createService := createBalanceSheet.NewCreateBalanceSheetService(createRepo)
	createHandler := handlerCreateBalanceSheet.NewCreateBalanceSheetHandler(createService)

	// Read
	readRepo := readBalanceSheet.NewReadBalanceSheetRepository(db)
	readService := readBalanceSheet.NewReadBalanceSheetService(readRepo)
	readHandler := handlerReadBalanceSheet.NewReadBalanceSheetService(readService)

	// Update
	updateRepo := updateBalanceSheet.NewUpdateBalanceSheetRepository(db)
	updateService := updateBalanceSheet.NewUpdateBalanceSheetService(updateRepo)
	updateHandler := handlerUpdateBalanceSheet.NewUpdateBalanceSheetHandler(updateService)

	groupRouter := routes.Group("/balance-sheets")

	groupRouter.POST("", createHandler.CreateBalanceSheet)

	groupRouter.GET("", readHandler.GetBalanceSheet)

	groupRouter.PUT("", updateHandler.UpdateBalanceSheet)
}
