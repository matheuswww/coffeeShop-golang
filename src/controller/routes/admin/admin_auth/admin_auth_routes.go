package admin_auth_routes

import (
	admin_auth_controller "matheuswww/coffeeShop-golang/src/controller/admin/admin_auth"
	"matheuswww/coffeeShop-golang/src/controller/routes/coockies"
	admin_auth_repository "matheuswww/coffeeShop-golang/src/model/admin/admin_auth/repository"
	admin_auth_service "matheuswww/coffeeShop-golang/src/model/admin/admin_auth/service"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func InitAdminAuthRoutes(r *gin.RouterGroup) {
	userController := initAdminAuthController()
	authGroup := r.Group("/admin/auth")
	authGroup.Use(sessions.Sessions("auth", coockies.Store()))

	authGroup.POST("/signIn", userController.SignIn)
}

func initAdminAuthController() admin_auth_controller.AdminAuthControllerInterface {
	adminAuthRepository := admin_auth_repository.NewAdminAuthRepository()
	adminAuthService := admin_auth_service.NewAuthDomainService(adminAuthRepository)
	adminAuthController := admin_auth_controller.NewAdminAuthInterface(adminAuthService)
	return adminAuthController
}
