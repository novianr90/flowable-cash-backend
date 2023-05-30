package main

import (
	"flowable-cash-backend/configs"
	"flowable-cash-backend/internal/sorting"
	"log"
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

	SetupInternalJob(db)
}

func SetupRouter(db *gorm.DB) *gin.Engine {

	router := gin.Default()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	routes.InitTransactionRoutes(db, router)

	return router
}

func SetupInternalJob(db *gorm.DB) {
	sort := sorting.NewSortingInternal(db)

	err := sort.SortTransaction()

	if err != nil {
		log.Fatal("error while sorting transaction", err)
	}
}
