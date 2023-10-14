package main

import (
	"database/sql"
	"log"
	"matheuswww/coffeeShop-golang/src/configuration/logger"
	"matheuswww/coffeeShop-golang/src/configuration/mysql"
	"os"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/joho/godotenv"
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		logger.Error("Error loading .env file", err)
		log.Fatal("Error loading .env file")
	}
}

func loadCors() *cors.Config {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{os.Getenv("CORSORIGIN")}
	config.AllowMethods = []string{"POST", "GET"}
	config.AllowHeaders = []string{"Origin", "X-Requested-With", "Content-Type", "Accept", "Cookie"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowCredentials = true
	return &config
}

func loadMysql() *sql.DB {
	host := os.Getenv("MYSQL_HOST")
	name := os.Getenv("MYSQL_NAME")
	password := os.Getenv("MYSQL_PASSWORD")
	port,err := strconv.Atoi(os.Getenv("MYSQL_PORT"))
	if err != nil {
		logger.Error("Error loading database", err)
		log.Fatal("invalid port")
	}
	db, err := mysql.NewMysql(host, name, password, port).NewMysqlConnection()
	if err != nil {
		logger.Error("Error loading database", err)
		log.Fatal("Error connecting")
	}
	return db
}