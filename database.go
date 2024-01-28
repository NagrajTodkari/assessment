package main

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB *gorm.DB
)

func InitConnectionsForJobs() {

	err := OpenDatabaseConnection()
	if err != nil {
		log.Fatal("DB error: ", err)
	}

}

func OpenDatabaseConnection() error {
	var err error

	logMod := logger.Info

	config := &gorm.Config{}
	config.Logger = logger.Default.LogMode(logMod)

	DB, err = gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), config)

	if err != nil {
		log.Fatal("DB error: ", err)
		return err
	}

	DB.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)
	DB.Exec(`CREATE EXTENSION IF NOT EXISTS hstore;`)

	return err
}
