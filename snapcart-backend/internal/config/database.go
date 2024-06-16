package config

import (
	"log"
	"os"

	"github.com/ulunnuha-h/snapcart/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func OpenDatabase(){
	var db *gorm.DB
	var err error
	var dsn string

	if os.Getenv("APP_ENV") == "local"{
		dsn = os.Getenv("POSTGRES_LOCAL")
	} else {
		dsn = os.Getenv("POSTGRES_PROD")
	}
	
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	db.AutoMigrate(&model.Product{})

	DB = db
}