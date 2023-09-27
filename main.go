package main

import (
	"matheuswww/coffeeShop-golang/src/configuration/logger"
	"matheuswww/coffeeShop-golang/src/controller/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	logger.Info("About to start user application")
	loadEnv()
	router := gin.Default()
	corsConfig := loadCors()
	router.Use(cors.New(*corsConfig))
	routes.InitRoutes(&router.RouterGroup)
	if err := router.Run(":8080"); err != nil {
		logger.Error("Error to load router", err)
		panic("Error to init router")
	}
}