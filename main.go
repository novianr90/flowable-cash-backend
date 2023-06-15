package main

import (
	"flowable-cash-backend/configs"
	"os"

	"gorm.io/gorm"

	"flowable-cash-backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	db := configs.Connection()

	PORT := os.Getenv("PORT")

	app := SetupRouter(db)

	app.Run(":" + PORT)
}

func SetupRouter(db *gorm.DB) *gin.Engine {

	router := gin.Default()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	apiRouter := router.Group("/api")

	routes.InitTransactionRoutes(db, apiRouter)
	routes.InitSaleTransactionRoutes(db, apiRouter)
	routes.InitPurchaseRoutes(db, apiRouter)
	routes.InitBalanceSheetRoutes(db, apiRouter)

	return router
}
