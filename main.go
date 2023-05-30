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
	PORT := os.Getenv("PORT")

	var wg *sync.WaitGroup

	app := SetupRouter(wg)

	app.Run(":" + PORT)

	wg.Wait()
}

func SetupRouter(wg *sync.WaitGroup) *gin.Engine {
	db := configs.Connection()

	wg.Add(1)
	go SetupInternalJob(db)

	router := gin.Default()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	routes.InitTransactionRoutes(db, router)

	defer wg.Done()

	return router
}

func SetupInternalJob(db *gorm.DB) {
	sort := sorting.NewSortingInternal(db)

	err := sort.SortTransaction()

	if err != nil {
		log.Fatal("error while sorting transaction", err)
	}
}
