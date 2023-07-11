package configs

import (
	"flowable-cash-backend/models"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

func Connection() *gorm.DB {
	var (
		db  *gorm.DB
		err error

		mySql = os.Getenv("MYSQL_URL")
	)

	db, err = gorm.Open(mysql.Open(mySql), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	fmt.Println("connected to database")

	err = db.Debug().AutoMigrate(
		models.Transaction{},
		models.BalanceSheet{},
	)

	if err != nil {
		log.Fatal(err)
	}

	return db
}
