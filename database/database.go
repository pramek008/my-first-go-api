package database

import (
	"fmt"
	"log"
	"os"

	"github.com/pramek008/first-golang/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	
)

type DbInstance struct{
	Db *gorm.DB
}

var DB DbInstance

func ConnectDb(){
	dsn := fmt.Sprintf(
		"host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Jakarta",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to DB. \n", err)
		os.Exit(1)
	}

	log.Println("DB Connected")
	db.Logger = logger.Default.LogMode(logger.Info)
	
	log.Println("Running DB Migrations")
	db.AutoMigrate(&models.Fact{})

	DB = DbInstance{
		Db: db,
	}	


}

