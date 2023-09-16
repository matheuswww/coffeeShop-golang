package main

import (
	"log"
	"matheuswww/coffeeShop-golang/src/configuration/logger"
	"matheuswww/coffeeShop-golang/src/configuration/mysql"
	"matheuswww/coffeeShop-golang/src/controller/routes"
	user_controller "matheuswww/coffeeShop-golang/src/controller/user"
	user_repository "matheuswww/coffeeShop-golang/src/model/user/repository"
	user_service "matheuswww/coffeeShop-golang/src/model/user/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("About to start user application")
	err := godotenv.Load()
	if err != nil {
		logger.Error("Error loading .env file",err)
		log.Fatal("Error loading .env file")
	}
	database,err := mysql.NewMysqlConnection()
	if err != nil {
		logger.Error("Error loading database",err)
		log.Fatal("Error loading database")
	}
	userRepository := user_repository.NewUserRepository(database)
	userService := user_service.NewUserDomainService(userRepository)
	userController := user_controller.NewUserControllerInterface(userService)

	router := gin.Default()
	routes.InitRouter(&router.RouterGroup, userController)
	if err := router.Run(":8080"); err != nil {
		logger.Error("Error to load router",err)
		log.Fatal(err)
	}
}