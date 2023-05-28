package main

import (
	"flowable-cash-backend/models"
	"flowable-cash-backend/router"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func init() {
	connString := os.Getenv("DATABASE_URL")

	fmt.Println(connString)

	db, err = gorm.Open(postgres.Open(connString), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	fmt.Println("connected to database")

	db.Debug().AutoMigrate(models.Transaction{}, models.General{}, models.Ledger{})
}

func main() {
	PORT := os.Getenv("PORT")

	router.StartServer(db).Run(":" + PORT)
}
