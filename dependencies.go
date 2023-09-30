package main

import (
	"database/sql"
	"errors"
	"log"
	"matheuswww/coffeeShop-golang/src/configuration/logger"
	"matheuswww/coffeeShop-golang/src/configuration/mysql"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/joho/godotenv"
)

var (
	mode string
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		logger.Error("Error loading .env file", err)
		log.Fatal("Error loading .env file")
	}
	mode = os.Getenv("MODE")
}

func loadCors() *cors.Config {
	config := cors.DefaultConfig()
	if mode == "PROD" {
		return nil
	} else if mode == "DEV" {
		config.AllowOrigins = []string{os.Getenv("CORSORIGIN")}
		config.AllowMethods = []string{"POST", "GET"}
		config.AllowHeaders = []string{"Origin", "X-Requested-With", "Content-Type", "Accept", "Cookie"}
		config.ExposeHeaders = []string{"Content-Length"}
		config.AllowCredentials = true
		return &config
	}
	logger.Error("Error loading database", errors.New("invalid mode"))
	log.Fatal("Error loading database")
	return &config
}

func loadMysql() *sql.DB {
	mode := os.Getenv("MODE")
	if mode == "PROD" {
		return nil
	} else if mode == "DEV" {
		host := "172.17.0.3"
		name := "coffeeShop"
		password := "senha"
		port := 8080
		db, err := mysql.NewMysql(host, name, password, port).NewMysqlConnection()
		if err != nil {
			logger.Error("Error loading database", err)
			log.Fatal("Error loading database")
		}
		return db
	}
	logger.Error("Error loading database", errors.New("invalid mode"))
	log.Fatal("Error loading database")
	return nil
}