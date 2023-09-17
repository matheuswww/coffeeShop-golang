package main

import (
	"database/sql"
	"log"
	"matheuswww/coffeeShop-golang/src/configuration/logger"
	"matheuswww/coffeeShop-golang/src/configuration/mysql"
	"os"

	"github.com/joho/godotenv"
)

var (
	mode string
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		logger.Error("Error loading .env file",err)
		log.Fatal("Error loading .env file")
	}
	mode = os.Getenv("MODE")
}

func InitDatabase() (*sql.DB) {
	if(mode == "DEV") {
		newDB := mysql.NewMysql("172.17.0.3","coffeeShop","senha",8080)
		database,err := newDB.NewMysqlConnection()
		if err != nil {
			logger.Error("Error loading database",err)
			log.Fatal("Error loading database")
		}
		return database
	}
	return nil
}