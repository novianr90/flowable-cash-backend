package main

import (
	"flowable-cash-backend/configs"
	"os"

	"flowable-cash-backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	PORT := os.Getenv("PORT")

	app := SetupRouter()

	app.Run(":" + PORT)
}

func SetupRouter() *gin.Engine {
	db := configs.Connection()

	router := gin.Default()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	routes.InitTransactionRoutes(db, router)

	return router
}
