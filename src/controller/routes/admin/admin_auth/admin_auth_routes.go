package admin_auth_routes

import (
	admin_auth_controller "matheuswww/coffeeShop-golang/src/controller/admin/admin_auth"
	admin_auth_repository "matheuswww/coffeeShop-golang/src/model/admin/admin_auth/repository"
	admin_auth_service "matheuswww/coffeeShop-golang/src/model/admin/admin_auth/service"

	"github.com/gin-gonic/gin"
)

func InitAdminAuthRoutes(r *gin.RouterGroup) {
	userController := InitAdminAuthController()
	authGroup := r.Group("/admin/auth")

	authGroup.POST("/signIn", userController.SignIn)
}

func InitAdminAuthController() admin_auth_controller.AdminAuthControllerInterface {
	adminAuthRepository := admin_auth_repository.NewAdminAuthRepository()
	adminAuthService := admin_auth_service.NewAuthDomainService(adminAuthRepository)
	adminAuthController := admin_auth_controller.NewAdminAuthInterface(adminAuthService)
	return adminAuthController
}
