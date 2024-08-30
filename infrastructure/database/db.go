package database

import (
	"fmt"
	"log"
	"portfolio/todoApp/entity"
	"portfolio/todoApp/infrastructure/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func handleDatabaseConnection() (*gorm.DB, error) {
	appConfig := config.GetAppConfig()
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		appConfig.DBHost, appConfig.DBPort, appConfig.DBUser, appConfig.DBPassword, appConfig.DBName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("error occured while trying to connect to database", err)
	}

	db.Debug().AutoMigrate(entity.Todo{})
	return db, nil
}

func GetDatabaseInstance() *gorm.DB {
	db, err := handleDatabaseConnection()
	if err != nil {
		log.Panic(err)
	}
	return db
}
