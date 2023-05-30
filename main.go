package main

import (
	"flowable-cash-backend/configs"
	"flowable-cash-backend/internal/sorting"
	"log"
	"os"
	"sync"

	"gorm.io/gorm"

	"flowable-cash-backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db := configs.Connection()

	PORT := os.Getenv("PORT")

	var wg *sync.WaitGroup

	wg.Add(1)
	go SetupInternalJob(db, wg)

	app := SetupRouter(db)

	app.Run(":" + PORT)

	wg.Wait()
}

func SetupRouter(db *gorm.DB) *gin.Engine {

	router := gin.Default()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	routes.InitTransactionRoutes(db, router)

	return router
}

func SetupInternalJob(db *gorm.DB, wg *sync.WaitGroup) {
	defer wg.Done()

	sort := sorting.NewSortingInternal(db)

	err := sort.SortTransaction()

	if err != nil {
		log.Fatal("error while sorting transaction", err)
	}
}
