package main

import (
	"log"
	"matheuswww/coffeeShop-golang/src/configuration/logger"
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
	if mode == "DEV"{
		config.AllowOrigins = []string{os.Getenv("CORSORIGIN")}
		config.AllowMethods = []string{"POST", "GET"}
		config.AllowHeaders = []string{"Origin", "X-Requested-With", "Content-Type", "Accept","Cookie"}
		config.ExposeHeaders = []string{"Content-Length"}
		config.AllowCredentials = true
		return &config
	}
	return &config
}