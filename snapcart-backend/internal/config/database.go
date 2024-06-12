package config

import (
	"fmt"
	"log"
	"os"

	"github.com/ulunnuha-h/snapcart/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenDatabase() *gorm.DB{
	var db *gorm.DB
	var err error

	dsn := os.Getenv("POSTGRES_URI")
	fmt.Print(dsn)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	db.AutoMigrate(&model.Product{})

	return db
}