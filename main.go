package main

import (
	"flowable-cash-backend/configs"
	"flowable-cash-backend/internal/sorting"
	"fmt"
	"log"
	"os"
	"sync"

	"gorm.io/gorm"

	"flowable-cash-backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	var wg *sync.WaitGroup

	db := configs.Connection()

	PORT := os.Getenv("PORT")

	app := SetupRouter(db)

	SetupInternalJob(db)

	wg.Add(1)
	go func() {
		defer wg.Done()
		app.Run(":" + PORT)
	}()

	wg.Wait()
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

	result, err := sort.SortTransaction()

	fmt.Println("Sorting is done")

	fmt.Println("Transactions size is", len(result))

	if err != nil {
		log.Fatal("error while sorting transaction", err)
	}
}
