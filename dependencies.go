package main

import (
	"log"
	"matheuswww/coffeeShop-golang/src/configuration/logger"
	"os"

	"github.com/joho/godotenv"
)

var (
	mode string
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		logger.Error("Error loading .env file", err)
		log.Fatal("Error loading .env file")
	}
	mode = os.Getenv("MODE")
}
