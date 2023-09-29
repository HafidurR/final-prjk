package config

import (
	"fmt"
	"log"
	"os"

	"github.com/HafidurR/final-prjk/src/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func Connect() *gorm.DB {
	// db, err = gorm.Open(mysql.Open("root:@tcp(localhost:3306)/db_final_prjk?parseTime=true"))

	dsn := fmt.Sprintf("%v:@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
				os.Getenv("DB_USERNAME"),
				// os.Getenv("DB_PASSWORD"),
				os.Getenv("DB_HOST"),
				os.Getenv("DB_PORT"),
				os.Getenv("DB_NAME"))
	
	fmt.Println("ya error cok", os.Getenv("DB_NAME"))	
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("connection failed", err)
	} else {
		log.Println("connection succeeded")
	}

	db.AutoMigrate(&models.Author{}, &models.Book{})
	return db
}
