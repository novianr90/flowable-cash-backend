package main

import (
	"flowable-cash-backend/models"
	"flowable-cash-backend/router"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func init() {

	var (
		username = os.Getenv("DATABASE_USERNAME")
		password = os.Getenv("DATABASE_PASSWORD")
		host     = os.Getenv("DATABASE_HOST")
		port     = os.Getenv("DATABASE_PORT")
		database = os.Getenv("DATABASE_NAME")
	)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True&loc=Local&tls=true", username, password, host, port, database)

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

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
