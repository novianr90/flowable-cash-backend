package main

import (
	"flowable-cash-backend/configs"
	"flowable-cash-backend/usecase"
	"log"
	"os"

	"gorm.io/gorm"

	"flowable-cash-backend/routes"

	"github.com/gin-gonic/gin"

	"github.com/robfig/cron/v3"
)

var (
	c = cron.New()
)

func main() {

	db := configs.Connection()

	PORT := os.Getenv("PORT")

	app := SetupRouter(db)

	go PostingPenjualanScheduler(db)

	app.Run(":" + PORT)
}

func PostingPenjualanScheduler(db *gorm.DB) {

	useCaseService := usecase.NewUseCaseService(db)

	_, err := c.AddFunc("@every 5m", func() {
		err := useCaseService.PostingPenjualan()
		if err != nil {
			log.Println("Error when posting:", err)
		}
	})

	if err != nil {
		log.Println("Error when do job:", err)
	}

	_, err = c.AddFunc("@every 5m", func() {
		err := useCaseService.PostingPembelian()
		if err != nil {
			log.Println("Error when posting:", err)
		}
	})

	if err != nil {
		log.Println("Error when do job:", err)
	}

	c.Start()

	select {}
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
	routes.InitPostingRoutes(db, apiRouter)

	return router
}
