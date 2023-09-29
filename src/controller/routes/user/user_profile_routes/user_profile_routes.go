package user_profile_routes

import (
	sessionCookie "matheuswww/coffeeShop-golang/src/controller/routes/cookies"
	user_profile_controller "matheuswww/coffeeShop-golang/src/controller/user/user_profile"
	user_profile_repository "matheuswww/coffeeShop-golang/src/model/user/user_profile/repository"
	user_profile_service "matheuswww/coffeeShop-golang/src/model/user/user_profile/service"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func InitUserProfileRoutes(r *gin.RouterGroup) {
	userController := initUserProfileController()
	authGroup := r.Group("/userProfile")
	authGroup.Use(sessions.Sessions("auth", sessionCookie.Store()))

	authGroup.POST("/addToCart", userController.AddToCart)
	authGroup.GET("/profile", userController.GetProfile)
}

func initUserProfileController() user_profile_controller.UserProfileController {
	userProfileRepository := user_profile_repository.NewUserProfileRepository()
	userProfileService := user_profile_service.NewUserProfileService(userProfileRepository)
	userProfileController := user_profile_controller.NewUserProfileController(userProfileService)
	return userProfileController
}
