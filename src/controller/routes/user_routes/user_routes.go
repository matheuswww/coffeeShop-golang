package user_routes

import (
	"database/sql"
	user_controller "matheuswww/coffeeShop-golang/src/controller/user"
	user_repository "matheuswww/coffeeShop-golang/src/model/user/repository"
	user_service "matheuswww/coffeeShop-golang/src/model/user/service"
	"os"

	"github.com/gin-gonic/gin"
)

func InitUserRoutes(r *gin.RouterGroup,database *sql.DB) {
	userController := initUserController(database)
	r.POST("/signUp",userController.SignUp)
}

func initUserController(database *sql.DB) user_controller.UserControllerInterface {
	if(os.Getenv("MODE") == "DEV") {
		userRepository := user_repository.NewUserRepository(database)
		userService := user_service.NewUserDomainService(userRepository)
		userController := user_controller.NewUserControllerInterface(userService)
		return userController
	}
	return nil
}