package user_routes

import (
	"database/sql"
	"matheuswww/coffeeShop-golang/src/controller/routes/coockies"
	user_auth_controller "matheuswww/coffeeShop-golang/src/controller/user/user_auth"
	user_auth_repository "matheuswww/coffeeShop-golang/src/model/user/user_auth/repository"
	user_auth_service "matheuswww/coffeeShop-golang/src/model/user/user_auth/service"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func InitUserAuthRoutes(r *gin.RouterGroup,database *sql.DB) {
	userController := initUserAuthController(database)
	authGroup := r.Group("/auth")
	authGroup.Use(sessions.Sessions("auth",coockies.Store()))

	authGroup.POST("/signUp",userController.SignUp)
	authGroup.POST("/signIn",userController.SignIn)
	authGroup.POST("/email",userController.SendAuthEmail)
	authGroup.GET("/email/:token",userController.AuthEmail)
}

func initUserAuthController(database *sql.DB) user_auth_controller.UserAuthControllerInterface {
	if(os.Getenv("MODE") == "DEV") {
		userAuthRepository := user_auth_repository.NewUserAuthRepository(database)
		userAuthService := user_auth_service.NewUserAuthDomainService(userAuthRepository)
		userAuthController := user_auth_controller.NewUserAuthControllerInterface(userAuthService)
		return userAuthController
	}
	return nil
}