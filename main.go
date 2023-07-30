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

	go PostingUseCase(db)

	app.Run(":" + PORT)
}

func PostingUseCase(db *gorm.DB) {

	useCaseService := usecase.NewUseCaseService(db)

	// Pemasukkan
	_, err := c.AddFunc("@every 5m", func() {
		err := useCaseService.PostingPemasukkan()
		if err != nil {
			log.Println("Error when posting:", err)
		}
	})

	if err != nil {
		log.Println("Error when do job:", err)
	}

	// Bahan Baku
	_, err = c.AddFunc("@every 6m", func() {
		err := useCaseService.PostingBahanBaku()
		if err != nil {
			log.Println("Error when posting:", err)
		}
	})

	if err != nil {
		log.Println("Error when do job:", err)
	}

	// Barang Dagang
	_, err = c.AddFunc("@every 7m", func() {
		err := useCaseService.PostingBarangDagang()
		if err != nil {
			log.Println("Error when posting:", err)
		}
	})

	if err != nil {
		log.Println("Error when do job:", err)
	}

	// Bahan Tambahan
	_, err = c.AddFunc("@every 8m", func() {
		err := useCaseService.PostingBahanTambahan()
		if err != nil {
			log.Println("Error when posting:", err)
		}
	})

	if err != nil {
		log.Println("Error when do job:", err)
	}

	// Peralatan
	_, err = c.AddFunc("@every 9m", func() {
		err := useCaseService.PostingPeralatan()
		if err != nil {
			log.Println("Error when posting:", err)
		}
	})

	if err != nil {
		log.Println("Error when do job:", err)
	}

	// Hutang
	_, err = c.AddFunc("@every 10m", func() {
		err := useCaseService.PostingBayarHutang()
		if err != nil {
			log.Println("Error when posting:", err)
		}
	})

	if err != nil {
		log.Println("Error when do job:", err)
	}

	// Piutang
	_, err = c.AddFunc("@every 11m", func() {
		err := useCaseService.PostingBayarPiutang()
		if err != nil {
			log.Println("Error when posting:", err)
		}
	})

	if err != nil {
		log.Println("Error when do job:", err)
	}

	// Biaya
	_, err = c.AddFunc("@every 12m", func() {
		err := useCaseService.PostingBiayaBiaya()
		if err != nil {
			log.Println("Error when posting:", err)
		}
	})

	if err != nil {
		log.Println("Error when do job:", err)
	}

	// Currently not implemented, still search how to balance if add this

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
