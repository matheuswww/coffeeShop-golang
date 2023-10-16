package admin_auth_routes

import (
	"database/sql"
	admin_auth_controller "matheuswww/coffeeShop-golang/src/controller/admin/admin_auth"
	sessionCookie "matheuswww/coffeeShop-golang/src/controller/routes/cookies"
	admin_auth_repository "matheuswww/coffeeShop-golang/src/model/admin/admin_auth/repository"
	admin_auth_service "matheuswww/coffeeShop-golang/src/model/admin/admin_auth/service"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func InitAdminAuthRoutes(r *gin.RouterGroup, database *sql.DB) {
	userController := initAdminAuthController(database)
	authGroup := r.Group("/admin/auth")
	authGroup.Use(sessions.Sessions("auth", sessionCookie.Store()))

	authGroup.POST("/signIn", userController.SignIn)
	authGroup.POST("/signUp", userController.CreateAdmin)
}

func initAdminAuthController(database *sql.DB) admin_auth_controller.AdminAuthControllerInterface {
	adminAuthRepository := admin_auth_repository.NewAdminAuthRepository(database)
	adminAuthService := admin_auth_service.NewAuthDomainService(adminAuthRepository)
	adminAuthController := admin_auth_controller.NewAdminAuthInterface(adminAuthService)
	return adminAuthController
}
